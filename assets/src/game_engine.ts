import * as ex from "excalibur"
import { PointerScope } from "excalibur/dist/Input/Index"
import * as msgs from "./messages"
import * as res from "./resources"
import { Planet } from "./planet"
import { Player } from "./player"
import { Cursor } from "./cursor"
import { getCookie } from "./utils"
import { ScreenSelector } from "./screen_selector"

const watchMoveInterval = 200

export class GameEngine extends ex.Engine {
    players: Map<string, Player>
    localPlayers: Set<string>
    currentPlayer: string
    myTurn: boolean
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
        this.myTurn = false
        this.cursor = new Cursor(this)
        this.currentPlayer = ""

        this.label = new ex.Label('Current player:', 0, 40)
        this.label.color = ex.Color.White
        this.label.fontFamily = 'Arial'
        this.label.fontSize = 40
        this.label.fontUnit = ex.FontUnit.Px
        this.label.textAlign = ex.TextAlign.Left

        this.timer = new ex.Timer({
            interval: watchMoveInterval,
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
                this.ss.restoreToWelcomeScreen()
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
                this.ss.restoreToWelcomeScreen()
            })
        })
    }

    enableCursor() {
        this.add(this.cursor)
        this.cursor.setZIndex(100)
    }

    disableCursor() {
        this.remove(this.cursor)
    }

    takeTurn(currentPlayer: string) {
        this.myTurn = true
        this.enableCursor()
        this.players.get(currentPlayer)?.activate()
        this.currentPlayer = currentPlayer
        this.label.color = ex.Color.Green
    }

    pollTurn() {
        var gameID = getCookie("game_id")

        if (!this.myTurn) {
            $.post("/game/poll_turn", { "game_id": gameID }, (data: msgs.PollTurn) => {
                this.label.text = `Current player: ${data.CurrentPlayer}`
                if (this.localPlayers.has(data.CurrentPlayer)) {
                    this.takeTurn(data.CurrentPlayer)
                } else {
                    var currentPlayer = this.players.get(data.CurrentPlayer)
                    currentPlayer?.setDestination(data.CurrentPlayerAlpha)
                    this.label.color = ex.Color.White
                }
            }, "json").fail(() => {
                this.ss.restoreToWelcomeScreen()
            })
        }
    }

    onPreUpdate() {
        if (this.input.keyboard.isHeld(ex.Input.Keys.Esc))
            this.ss.restoreToWelcomeScreen()
    }
}