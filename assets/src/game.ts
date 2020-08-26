import * as ex from "excalibur"
import { getCanvas } from "./utils"

export class Game {
    game: ex.Engine
    canvas: HTMLCanvasElement

    constructor() {
        this.canvas = getCanvas("game")

        this.game = new ex.Engine({
            canvasElement: this.canvas,
        })

        this.canvas.hidden = true
    }
}
