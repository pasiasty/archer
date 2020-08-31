import * as ex from "excalibur"
import { line } from "excalibur/dist/Util/DrawUtil"

export class Cursor extends ex.Actor {
    shootStartPoint: ex.Vector
    shootEndPoint: ex.Vector
    engine: ex.Engine

    constructor(engine: ex.Engine) {
        super()

        this.engine = engine

        this.shootStartPoint = new ex.Vector(-1, -1)
        this.shootEndPoint = new ex.Vector(-1, -1)

        this.engine.input.pointers.primary.on('down', (evt: ex.Input.PointerDownEvent) => {
            this.shootStartPoint = evt.pos
        })

        this.engine.input.pointers.primary.on('up', (_: ex.Input.PointerUpEvent) => {
            this.shootStartPoint = new ex.Vector(-1, -1)
        })

        this.engine.input.pointers.primary.on("move", (evt: ex.Input.PointerMoveEvent) => {
            this.shootEndPoint = evt.pos
        })
    }

    public onPostDraw(ctx: CanvasRenderingContext2D, delta: number) {
        if (this.shootStartPoint.x != -1)
            line(ctx, ex.Color.Green, this.shootStartPoint.x, this.shootStartPoint.y, this.shootEndPoint.x, this.shootEndPoint.y, 3)
    }
}