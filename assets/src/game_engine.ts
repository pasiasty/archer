import * as ex from "excalibur"
import { PointerScope } from "excalibur/dist/Input/Index"
import * as msgs from "./messages"
import * as res from "./resources"
import { Planet } from "./planet"
import { Player } from "./player"
import { Cursor } from "./cursor"
import { getCookie } from "./utils"

export class GameEngine extends ex.Engine {
    players: Map<string, Player>
    localPlayers: Set<string>
    currentPlayer: string
    myTurn: boolean
    planets: Map<number, Planet>
    cursor: Cursor

    constructor() {
        super({
            canvasElementId: "game_screen",
            pointerScope: PointerScope.Canvas,
            displayMode: ex.DisplayMode.Fixed,
            width: 1920,
            height: 1080,
            backgroundColor: ex.Color.Black,
            suppressPlayButton: true,
        })

        this.players = new Map<string, Player>()
        this.localPlayers = new Set<string>()
        this.planets = new Map<number, Planet>()
        this.myTurn = false
        this.cursor = new Cursor(this)
        this.currentPlayer = ""
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
                    this.players.set(p.Name, new Player(planet, p.Alpha, p.ColorIdx))
                }
            }, "json").fail(() => {
                alert("failed to get world")
            })

            $.post("/preparation/list_users", { "game_id": gameID }, (data: msgs.UsersList) => {
                for (let u of data.Users) {
                    if (`"${u.Username}"` === username) {
                        for (let p of u.Players) {
                            this.localPlayers.add(p)
                        }
                    }
                }
            }, "json").fail(() => {
                alert("failed to get users_list")
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

    onPreUpdate() {
        var gameID = getCookie("game_id")
        var userID = getCookie("user_id")

        if (!this.myTurn) {
            $.post("/game/poll_turn", { "game_id": gameID }, (data: msgs.PollTurn) => {
                if (this.localPlayers.has(data.CurrentPlayer)) {
                    this.myTurn = true
                    this.enableCursor()
                    this.players.get(data.CurrentPlayer)?.activate()
                    this.currentPlayer = data.CurrentPlayer
                } else {
                    var currentPlayer = this.players.get(data.CurrentPlayer)
                    if (currentPlayer != null) {
                        currentPlayer.rotation = data.CurrentPlayerAlpha
                    }
                }
            }, "json").fail(() => {
                alert("failed to post poll_turn")
            })
        } else {
            var newAlpha = this.players.get(this.currentPlayer)?.rotation
            $.post("/game/move", { "game_id": gameID, "user_id": userID, "username": this.currentPlayer, "new_alpha": newAlpha }, null, "json").fail(() => {
                alert("failed to post move")
            })
        }
    }
}