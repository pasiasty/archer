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
    emitter: ex.ParticleEmitter

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

        this.emitter = new ex.ParticleEmitter();
        this.emitter.emitterType = ex.EmitterType.Rectangle;
        this.emitter.radius = 5;
        this.emitter.minVel = 100;
        this.emitter.maxVel = 200;
        this.emitter.minAngle = 3.2;
        this.emitter.maxAngle = 3.6;
        this.emitter.isEmitting = true;
        this.emitter.emitRate = 50;
        this.emitter.opacity = 0.5;
        this.emitter.fadeFlag = true;
        this.emitter.particleLife = 400;
        this.emitter.maxSize = 10;
        this.emitter.minSize = 1;
        this.emitter.startSize = 0;
        this.emitter.endSize = 0;
        this.emitter.acceleration = new ex.Vector(800, 0);
        this.emitter.beginColor = ex.Color.Orange;
        this.emitter.endColor = ex.Color.Yellow;

        this.game.add(this)
        this.game.add(this.emitter)
        this.setZIndex(30)
        this.emitter.setZIndex(10)
    }

    public onPostUpdate(engine: ex.Engine, delta: number) {
        if (!this.keepFlying)
            return
        if (this.sampleIdx == this.trajectory.ArrowStates.length) {
            this.emitter.kill()
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

            this.emitter.pos = this.pos.sub(new ex.Vector(Math.cos(this.rotation) * 55, Math.sin(this.rotation) * 55))
            this.emitter.rotation = this.rotation

            this.sampleIdx++
        }
    }
}