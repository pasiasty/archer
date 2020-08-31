import * as msgs from "./messages"
import * as ex from "excalibur"
import { getPlanetTexture } from "./resources"
import { Vector } from "excalibur"

export class Planet extends ex.Actor {
    radius: number
    planetID: number

    constructor(p: msgs.Planet) {
        super({ x: p.Location.X, y: p.Location.Y })
        this.radius = p.Radius
        this.planetID = p.PlanetID

        const sprite = getPlanetTexture(p.ResourceID).asSprite().clone()
        sprite.scale = new Vector(2 * p.Radius / sprite.width, 2 * p.Radius / sprite.height)
        this.addDrawing(sprite)
    }
}