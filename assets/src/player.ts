import * as ex from "excalibur"
import { Planet } from "./planet"
import * as res from "./resources"

const omega = 0.0025

export class Player extends ex.Actor {
    colorID: number
    activated: boolean

    constructor(p: Planet, alpha: number, colorID: number) {
        super()

        this.activated = false
        this.colorID = colorID
        var playerSprite = res.Images.player.asSprite().clone()
        playerSprite.scale = new ex.Vector(0.1, 0.1)
        playerSprite.offset = new ex.Vector(0, 25 + p.radius)
        this.addDrawing(playerSprite)
        this.rotation = alpha
        p.add(this)
    }

    public update(engine: ex.Engine, delta: number) {
        if (!this.activated) {
            return
        }

        if (engine.input.keyboard.isHeld(ex.Input.Keys.Right))
            this.rotation += delta * omega
        if (engine.input.keyboard.isHeld(ex.Input.Keys.Left))
            this.rotation -= delta * omega
    }

    public activate() {
        this.activated = true
    }
}