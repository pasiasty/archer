import * as ex from "excalibur"

export class Game {
    game: ex.Engine

    constructor() {
        this.game = new ex.Engine({
            canvasElementId: "game",
        })
    }
}
