import * as ex from "excalibur"
import { Screen } from "./screen"
import { ScreenSelector } from "./screen_selector"
import { PointerScope } from "excalibur/dist/Input/Index"

export class GameScreen extends Screen {
    game: ex.Engine

    constructor(ss: ScreenSelector) {
        super("game", ss)

        this.game = new ex.Engine({
            pointerScope: PointerScope.Canvas
        })

        this.disable()
    }

    enable() {
        super.enable()
        this.game.start()
    }

    disable() {
        super.disable()
        this.game.stop()
    }
}
