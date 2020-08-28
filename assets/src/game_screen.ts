import * as ex from "excalibur"
import { Screen } from "./screen"
import { ScreenSelector } from "./screen_selector"
import * as msgs from "./messages"
import { PointerScope } from "excalibur/dist/Input/Index"
import { DrawUtil } from "excalibur/dist/Util/Index"

class Planet extends ex.Actor {
    p: msgs.Planet
    constructor(p: msgs.Planet) {
        super({ x: p.Location.X, y: p.Location.Y })
        this.p = p
    }

    public onPostDraw(ctx: CanvasRenderingContext2D, delta: number) {
        DrawUtil.circle(ctx, this.p.Location.X, this.p.Location.Y, this.p.Radius, ex.Color.White, ex.Color.White)
    }
}

export class GameScreen extends Screen {
    game: ex.Engine

    constructor(ss: ScreenSelector) {
        super("game_screen", ss)

        this.game = new ex.Engine({
            pointerScope: PointerScope.Canvas,
            displayMode: ex.DisplayMode.FullScreen,
            width: 1680,
            height: 1050,
            backgroundColor: ex.Color.Blue,
        })

        this.disable()
    }

    enable() {
        super.enable()
        this.game.start()

        $.post("/game/get_world", (data: msgs.PublicWorld) => {
            for (let p of data.Planets) {
                console.log("planet", p.Location.X, p.Location.Y, p.Radius)
                this.game.add(new Planet(p))
            }
        }, "json").fail(() => {
            alert("failed to get world")
        })
    }

    disable() {
        super.disable()
        this.game.stop()
    }
}
