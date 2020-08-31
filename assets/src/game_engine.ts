import * as ex from "excalibur"
import { PointerScope } from "excalibur/dist/Input/Index"
import * as msgs from "./messages"
import * as res from "./resources"
import { Planet } from "./planet"
import { Player } from "./player"
import { Cursor } from "./cursor"

export class GameEngine extends ex.Engine {
    players: Map<string, Player>

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
    }

    run() {
        console.log(this.drawWidth, this.canvasWidth)
        this.start(res.loader).then(() => {
            var background = new ex.Actor(this.halfDrawWidth, this.halfDrawHeight)
            background.addDrawing(res.Images.sky)
            this.add(background)

            $.post("/game/get_world", (data: msgs.PublicWorld) => {
                for (let p of data.Planets) {
                    console.log("planet", p.Location.X, p.Location.Y, p.Radius)
                    var newPlanet = new Planet(p)
                    new Player(newPlanet, 0.13)
                    this.add(newPlanet)
                }
            }, "json").fail(() => {
                alert("failed to get world")
            })

            var cursor = new Cursor(this)
            this.add(cursor)
            cursor.setZIndex(100)
        })
    }

    onPostDraw() {

    }
}