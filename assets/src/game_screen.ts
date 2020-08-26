import * as ex from "excalibur"
import { Screen } from "./screen"

export class GameScreen extends Screen {
    constructor() {
        super("game")
        this.disable()
    }
}
