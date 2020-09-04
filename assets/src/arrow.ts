import * as ex from "excalibur"
import { Trajectory } from "./messages";
import { Consts } from "./constants"
import * as res from "./resources"
import { GameEngine } from "./game_engine";

export class Arrow extends ex.Actor {
    trajectory: Trajectory
    sampleIdx: number
    game: ex.Engine
    callback: (game: ex.Engine, collidedWith: string) => void
    accDelta: number
    keepFlying: boolean

    constructor(game: ex.Engine, callback: (game: ex.Engine, collidedWith: string) => void, trajectory: Trajectory, color: ex.Color) {
        super()
        this.trajectory = trajectory
        this.sampleIdx = 0
        this.accDelta = 0
        this.keepFlying = true

        var arrowSprite = res.Images.arrow.asSprite().clone()
        arrowSprite.scale = new ex.Vector(0.02, 0.02)
        arrowSprite.offset = new ex.Vector(40, 22)
        arrowSprite.rotation = Math.PI
        arrowSprite.colorize(color)
        this.addDrawing(arrowSprite)

        this.game = game
        this.callback = callback

        this.game.add(this)
    }

    public onPostUpdate(engine: ex.Engine, delta: number) {
        if (!this.keepFlying)
            return
        if (this.sampleIdx == this.trajectory.ArrowStates.length) {
            this.keepFlying = false
            if (this.trajectory.CollidedWith != "planet") {
                this.game.remove(this)
            }
            this.callback(this.game, this.trajectory.CollidedWith)
            return
        }

        this.accDelta += delta

        if (this.accDelta > Consts.trajectoryInterval) {
            this.accDelta -= Consts.trajectoryInterval

            var arrowState = this.trajectory.ArrowStates[this.sampleIdx]
            this.pos = new ex.Vector(arrowState.Position.X, arrowState.Position.Y)
            this.rotation = arrowState.Orientation
            this.sampleIdx++
        }
    }
}