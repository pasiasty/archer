import * as msgs from "./messages"
import * as ex from "excalibur"
import { getPlanetTexture } from "./resources"
import { Vector } from "excalibur"

export class Planet extends ex.Actor {
    constructor(p: msgs.Planet) {
        super({ x: p.Location.X, y: p.Location.Y })

        const sprite = getPlanetTexture(p.ResourceID).asSprite().clone()
        sprite.scale = new Vector(2 * p.Radius / sprite.width, 2 * p.Radius / sprite.height)
        this.addDrawing(sprite)
    }
}