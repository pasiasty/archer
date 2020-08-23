import * as ex from "excalibur/dist/excalibur.min.js"

export class MainMenu extends ex.Scene {
    public onInitialize(engine: ex.Engine) {
        const paddle = new ex.Actor(150, 300, 200, 20)
        paddle.color = ex.Color.Chartreuse
        this.add(paddle)
    }
}
