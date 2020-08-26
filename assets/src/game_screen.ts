import * as ex from "excalibur"
import { Screen } from "./screen"
import { ScreenSelector } from "./screen_selector"

export class GameScreen extends Screen {
    constructor(ss: ScreenSelector) {
        super("game", ss)

        // var game = new ex.Engine({})
        // game.start()

        this.disable()
    }
}
