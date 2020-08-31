import * as ex from "excalibur"
import { Planet } from "./planet"
import * as res from "./resources"

const omega = 0.0025

export class Player extends ex.Actor {
    alpha: number
    constructor(p: Planet, alpha: number) {
        super()

        this.alpha = alpha
        var playerSprite = res.Images.player.asSprite().clone()
        playerSprite.scale = new ex.Vector(0.1, 0.1)
        playerSprite.offset = new ex.Vector(0, 25 + p.radius)
        this.addDrawing(playerSprite)
        this.rotation = this.alpha
        p.add(this)
    }

    public update(engine: ex.Engine, delta: number) {
        if (engine.input.keyboard.isHeld(ex.Input.Keys.Right))
            this.alpha += delta * omega
        if (engine.input.keyboard.isHeld(ex.Input.Keys.Left))
            this.alpha -= delta * omega

        this.rotation = this.alpha
    }
}