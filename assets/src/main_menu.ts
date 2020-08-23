import * as ex from "excalibur"

export class MainMenu extends ex.Scene {
    public onInitialize(engine: any) {
        const paddle = new ex.Actor(150, engine.drawHeight - 40, 200, 20)
        paddle.color = ex.Color.Chartreuse
        this.add(paddle)
    }
}
