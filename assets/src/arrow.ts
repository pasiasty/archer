import * as ex from "excalibur"
import { Trajectory } from "./messages";
import { Consts } from "./constants"
import * as res from "./resources"

export class Arrow extends ex.Actor {
    trajectory: Trajectory
    timer: ex.Timer
    sampleIdx: number
    game: ex.Engine
    callback: (game: ex.Engine, collidedWith: string) => void

    constructor(game: ex.Engine, callback: (game: ex.Engine, collidedWith: string) => void, trajectory: Trajectory, color: ex.Color) {
        super()
        this.trajectory = trajectory
        this.sampleIdx = 0

        this.timer = new ex.Timer({
            interval: Consts.trajectoryInterval,
            repeats: true,
            fcn: () => { this.updatePosition() },
        })

        var arrowSprite = res.Images.arrow.asSprite().clone()
        arrowSprite.scale = new ex.Vector(0.02, 0.02)
        arrowSprite.offset = new ex.Vector(40, 22)
        arrowSprite.rotation = Math.PI
        arrowSprite.colorize(color)
        this.addDrawing(arrowSprite)

        this.game = game
        this.callback = callback

        this.game.add(this)
        this.game.add(this.timer)
    }

    updatePosition() {
        if (this.sampleIdx == this.trajectory.ArrowStates.length) {
            this.timer.pause()

            if (this.trajectory.CollidedWith == "planet") {
                var removalTimer = new ex.Timer({
                    interval: Consts.arrowRemovalDelay,
                    fcn: () => { this.game.remove(this) },
                })
                this.game.add(removalTimer)
            }
            else {
                this.game.remove(this)
            }
            this.callback(this.game, this.trajectory.CollidedWith)
            return
        }
        var arrowState = this.trajectory.ArrowStates[this.sampleIdx]
        this.pos = new ex.Vector(arrowState.Position.X, arrowState.Position.Y)
        this.rotation = arrowState.Orientation
        this.sampleIdx++
    }
}