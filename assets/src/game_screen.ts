import * as ex from "excalibur"
import { Screen } from "./screen"
import { ScreenSelector } from "./screen_selector"
import * as msgs from "./messages"
import { PointerScope } from "excalibur/dist/Input/Index"
import { DrawUtil } from "excalibur/dist/Util/Index"
import { GameEngine } from "./game_engine"

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
    game: GameEngine

    constructor(ss: ScreenSelector) {
        super("game_screen", ss)

        this.game = new GameEngine(ss)
        this.disable()
    }

    enable() {
        super.enable()
        this.game.run()
    }

    disable() {
        super.disable()
        this.game.stop()
    }
}
