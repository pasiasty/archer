import * as ex from "excalibur"
import { PointerScope } from "excalibur/dist/Input/Index"
import * as msgs from "./messages"
import * as res from "./resources"
import { Planet } from "./planet"
import { Player } from "./player"
import { Cursor } from "./cursor"
import { getCookie } from "./utils"
import { ScreenSelector } from "./screen_selector"
import { Consts } from "./constants"
import { Arrow } from "./arrow"

export class GameEngine extends ex.Engine {
    players: Map<string, Player>
    localPlayers: Set<string>
    currentPlayer: string
    planets: Map<number, Planet>
    cursor: Cursor
    label: ex.Label
    ss: ScreenSelector
    timer: ex.Timer

    constructor(ss: ScreenSelector) {
        super({
            canvasElementId: "game_screen",
            pointerScope: PointerScope.Canvas,
            displayMode: ex.DisplayMode.Fixed,
            width: 1920,
            height: 1080,
            backgroundColor: ex.Color.Black,
            suppressPlayButton: true,
        })

        this.ss = ss
        this.players = new Map<string, Player>()
        this.localPlayers = new Set<string>()
        this.planets = new Map<number, Planet>()
        this.cursor = new Cursor(this, this.performShot)
        this.currentPlayer = ""

        this.label = new ex.Label('Current player:', 0, 40)
        this.label.color = ex.Color.White
        this.label.fontFamily = 'Arial'
        this.label.fontSize = 40
        this.label.fontUnit = ex.FontUnit.Px
        this.label.textAlign = ex.TextAlign.Left

        this.timer = new ex.Timer({
            interval: Consts.watchMoveInterval,
            fcn: () => {
                this.pollTurn()
            },
            repeats: true,
        })

        this.add(this.label)
        this.label.setZIndex(50)
    }

    run() {
        this.start(res.loader).then(() => {
            var background = new ex.Actor(this.halfDrawWidth, this.halfDrawHeight)
            background.addDrawing(res.Images.sky)
            this.add(background)

            var gameID = getCookie("game_id")
            var username = getCookie("username")

            $.post("/game/get_world", { "game_id": gameID }, (data: msgs.PublicWorld) => {
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
                this.add(this.timer)
            }, "json").fail(() => {
                this.ss.restoreToWelcomeScreen(true)
            })
        })
    }

    performShot(self: ex.Engine, v: ex.Vector) {
        var g = self as GameEngine
        if (v.size > 50) {
            var gameID = getCookie("game_id")
            var userID = getCookie("user_id")

            var currentPlayer = g.players.get(g.currentPlayer)

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

    afterPlayingTrajectory(self: ex.Engine) {
        var g = self as GameEngine
        g.add(g.timer)
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
        this.remove(this.timer)
        this.enableCursor()
        this.players.get(currentPlayer)?.activate()
        this.currentPlayer = currentPlayer
        this.label.color = ex.Color.Green
    }

    endTurn(data: msgs.Trajectory) {
        new Arrow(this, this.afterPlayingTrajectory, data)
        this.disableCursor()
        this.players.get(this.currentPlayer)?.deactivate()
        this.currentPlayer = ""
        this.label.color = ex.Color.White
    }

    pollTurn() {
        var gameID = getCookie("game_id")

        $.post("/game/poll_turn", { "game_id": gameID }, (data: msgs.PollTurn) => {
            this.label.text = `Current player: ${data.CurrentPlayer}`
            if (this.localPlayers.has(data.CurrentPlayer)) {
                this.takeTurn(data.CurrentPlayer)
            } else {
                var currentPlayer = this.players.get(data.CurrentPlayer)
                currentPlayer?.setDestination(data.CurrentPlayerAlpha)
                this.label.color = ex.Color.White

                if (data.ShotPerformed) {
                    this.remove(this.timer)
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