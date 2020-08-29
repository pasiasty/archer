import * as ex from "excalibur"
import { PointerScope } from "excalibur/dist/Input/Index"
import * as msgs from "./messages"
import * as res from "./resources"
import { Planet } from "./planet"

export class GameEngine extends ex.Engine {
    constructor() {
        super({
            pointerScope: PointerScope.Canvas,
            displayMode: ex.DisplayMode.Fixed,
            width: 1920,
            height: 1080,
            backgroundColor: ex.Color.Black,
            suppressPlayButton: true,
        })
    }

    run() {
        this.start(res.loader).then(() => {
            var background = new ex.Actor(this.halfDrawWidth, this.halfDrawHeight)
            background.addDrawing(res.Images.sky)
            this.add(background)

            $.post("/game/get_world", (data: msgs.PublicWorld) => {
                for (let p of data.Planets) {
                    console.log("planet", p.Location.X, p.Location.Y, p.Radius)
                    this.add(new Planet(p))
                }
            }, "json").fail(() => {
                alert("failed to get world")
            })
        })
    }
}