import * as ex from "excalibur"
import { Consts } from "./constants"
import { circle } from "excalibur/dist/Util/DrawUtil"

export class Indicator extends ex.Actor {
    growing: boolean
    radius: number
    color: ex.Color

    constructor(color: ex.Color, y: number) {
        super(0, y)

        this.color = color.average(ex.Color.White)
        this.growing = true
        this.radius = Consts.indicatorMinRadius
    }

    public onPostUpdate(engine: ex.Engine, delta: number) {
        if (this.growing) {
            this.radius += delta * Consts.indicatorChangeVel
            if (this.radius >= Consts.indicatorMaxRadius) {
                this.radius = Consts.indicatorMaxRadius
                this.growing = false
            }
        } else {
            this.radius -= delta * Consts.indicatorChangeVel
            if (this.radius <= Consts.indicatorMinRadius) {
                this.radius = Consts.indicatorMinRadius
                this.growing = true
            }
        }
    }

    public onPostDraw(ctx: CanvasRenderingContext2D, delta: number) {
        circle(ctx, 0, 0, this.radius, this.color, this.color)
    }
}