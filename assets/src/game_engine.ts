import * as ex from "excalibur"
import { PointerScope, KeyEvent, Keys } from "excalibur/dist/Input/Index"
import * as msgs from "./messages"
import * as res from "./resources"
import { Planet } from "./planet"
import { Player } from "./player"
import { Cursor } from "./cursor"
import { getCookie, optimalViewport } from "./utils"
import { ScreenSelector } from "./screen_selector"
import { Consts } from "./constants"
import { Arrow } from "./arrow"
import { time } from "console"

export class GameEngine extends ex.Engine {
    players: Map<string, Player>
    localPlayers: Set<string>
    currentPlayer: string
    planets: Map<number, Planet>
    cursor: Cursor
    currentPlayerLabel: ex.Label
    ss: ScreenSelector
    watchMoveTimer: ex.Timer
    autoResizeOn: boolean
    shootTimeout: number
    shootTime: number
    shootTimeLabel: ex.Label
    shootingTimer: ex.Timer

    constructor(ss: ScreenSelector) {
        super({
            canvasElementId: "game_screen",
            pointerScope: PointerScope.Canvas,
            displayMode: ex.DisplayMode.Fixed,
            viewport: optimalViewport(),
            resolution: { width: 1920, height: 1080 },
            backgroundColor: ex.Color.Black,
            suppressPlayButton: true,
        })
        this.autoResizeOn = true
        window.addEventListener("resize", () => {
            if (this.autoResizeOn) {
                this.screen.viewport = optimalViewport()
                this.screen.applyResolutionAndViewport()
            }
        });

        this.input.keyboard.on('press', (evt: KeyEvent) => {
            if (evt.key == Keys.F) {
                this.screen.goFullScreen()
            }
        })

        this.isDebug = Boolean(Consts.enableDebug)

        this.ss = ss
        this.players = new Map<string, Player>()
        this.localPlayers = new Set<string>()
        this.planets = new Map<number, Planet>()
        this.cursor = new Cursor(this, this.performShot)
        this.currentPlayer = ""

        this.currentPlayerLabel = new ex.Label('Current player:', 0, 40)
        this.currentPlayerLabel.color = ex.Color.White
        this.currentPlayerLabel.fontFamily = 'Arial'
        this.currentPlayerLabel.fontSize = 40
        this.currentPlayerLabel.fontUnit = ex.FontUnit.Px
        this.currentPlayerLabel.textAlign = ex.TextAlign.Left

        this.shootTimeout = 0
        this.shootTime = 0
        this.shootTimeLabel = new ex.Label('', 1850, 40)
        this.shootTimeLabel.color = ex.Color.White
        this.shootTimeLabel.fontFamily = 'Arial'
        this.shootTimeLabel.fontSize = 40
        this.shootTimeLabel.fontUnit = ex.FontUnit.Px
        this.shootTimeLabel.textAlign = ex.TextAlign.Left

        this.watchMoveTimer = new ex.Timer({
            interval: Consts.watchMoveInterval,
            fcn: () => {
                this.pollTurn()
            },
            repeats: true,
        })

        this.shootingTimer = this.newShootingTimer()

        this.add(this.currentPlayerLabel)
        this.currentPlayerLabel.setZIndex(50)
        this.add(this.shootTimeLabel)
        this.shootTimeLabel.setZIndex(50)
    }

    newShootingTimer(): ex.Timer {
        this.shootTime = this.shootTimeout
        this.shootTimeLabel.text = this.shootTime.toString()
        return new ex.Timer({
            interval: 1000,
            fcn: () => {
                this.shootTime--

                if (this.shootTime == 0) {
                    this.remove(this.shootingTimer)
                    this.shootTimeLabel.text = ''
                    this.performShot(this, new ex.Vector(0, 0), false)
                } else {
                    this.shootTimeLabel.text = this.shootTime.toString()
                }
            },
            numberOfRepeats: this.shootTimeout,
            repeats: true,
        })
    }

    run() {
        this.start(res.loader).then(() => {
            var background = new ex.Actor(this.halfDrawWidth, this.halfDrawHeight)
            background.addDrawing(res.Images.sky)
            this.add(background)

            var gameID = getCookie("game_id")
            var username = getCookie("username")

            $.post("/game/get_world", { "game_id": gameID }, (data: msgs.PublicWorld) => {
                this.shootTimeout = data.WorldSettings.ShootTimeout

                for (let p of data.Planets) {
                    var newPlanet = new Planet(p)
                    this.planets.set(newPlanet.planetID, newPlanet)
                    this.add(newPlanet)
                }

                for (let p of data.Players) {
                    var planet = this.planets.get(p.PlanetID)
                    if (planet == null)
                        continue
                    this.players.set(p.Name, new Player(p.Name, planet, p.Alpha, p.ColorIdx, this.ss, this))
                }
            }, "json").fail(() => {
                this.ss.restoreToWelcomeScreen(true)
            })

            $.post("/preparation/list_users", { "game_id": gameID }, (data: msgs.UsersList) => {
                for (let u of data.Users) {
                    if (`"${u.Username}"` === username) {
                        for (let p of u.Players) {
                            this.localPlayers.add(p)
                        }
                    }
                }
                this.add(this.watchMoveTimer)
            }, "json").fail(() => {
                this.ss.restoreToWelcomeScreen(true)
            })
        })
    }

    performShot(self: ex.Engine, v: ex.Vector, checkSize = true) {
        var g = self as GameEngine

        if (!checkSize || v.size > 50) {
            g.remove(g.shootingTimer)
            g.shootTimeLabel.text = ''

            var gameID = getCookie("game_id")
            var userID = getCookie("user_id")

            var currentPlayer = g.players.get(g.currentPlayer)
            currentPlayer?.deactivate()
            g.disableCursor()
            g.currentPlayerLabel.color = ex.Color.White

            $.post("/game/shoot", {
                "game_id": gameID,
                "user_id": userID,
                "username": currentPlayer?.username,
                "new_alpha": currentPlayer?.rotation,
                "shot_x": v.x,
                "shot_y": v.y,
            }, (data: msgs.Trajectory) => {
                g.endTurn(data)
            }, "json").fail(() => {
                g.ss.restoreToWelcomeScreen(true)
            })
        }
    }

    afterPlayingTrajectory(self: ex.Engine, collidedWith: string) {
        var g = self as GameEngine
        g.players.get(collidedWith)?.killPlayer()
        g.add(g.watchMoveTimer)
    }

    enableCursor() {
        this.add(this.cursor)
        this.cursor.enabled = true
        this.cursor.setZIndex(100)
    }

    disableCursor() {
        this.remove(this.cursor)
        this.cursor.enabled = false
    }

    takeTurn(currentPlayer: string) {
        this.remove(this.watchMoveTimer)
        this.enableCursor()
        this.players.get(currentPlayer)?.activate()
        this.currentPlayer = currentPlayer
        this.currentPlayerLabel.color = ex.Color.Green

        if (this.shootTimeout != 0) {
            this.shootingTimer = this.newShootingTimer()
            this.add(this.shootingTimer)
        }
    }

    endTurn(data: msgs.Trajectory) {
        var currentPlayer = this.players.get(this.currentPlayer)
        if (currentPlayer != null) {
            new Arrow(this, this.afterPlayingTrajectory, data, currentPlayer.playerColor)
            this.currentPlayer = ""
        }
    }

    pollTurn() {
        var gameID = getCookie("game_id")

        $.post("/game/poll_turn", { "game_id": gameID }, (data: msgs.PollTurn) => {
            this.currentPlayerLabel.text = `Current player: ${data.CurrentPlayer}`
            if (this.localPlayers.has(data.CurrentPlayer)) {
                this.takeTurn(data.CurrentPlayer)
            } else {
                var currentPlayer = this.players.get(data.CurrentPlayer)
                this.currentPlayer = data.CurrentPlayer
                currentPlayer?.setDestination(data.CurrentPlayerAlpha)
                this.currentPlayerLabel.color = ex.Color.White

                if (data.ShotPerformed) {
                    this.remove(this.watchMoveTimer)
                    this.getTrajectory(this)
                }
            }
        }, "json").fail(() => {
            this.ss.restoreToWelcomeScreen(true)
        })
    }

    getTrajectory(self: GameEngine) {
        var gameID = getCookie("game_id")
        var userID = getCookie("user_id")
        var username = self.localPlayers.values().next().value

        $.post("/game/get_trajectory", {
            "game_id": gameID,
            "user_id": userID,
            "username": username,
        }, (data: msgs.Trajectory) => {
            self.endTurn(data)
        }, "json").fail(() => {
            self.ss.restoreToWelcomeScreen(true)
        })
    }

    onPreUpdate() {
        if (this.input.keyboard.isHeld(ex.Input.Keys.Esc))
            this.ss.restoreToWelcomeScreen(true)
    }
}