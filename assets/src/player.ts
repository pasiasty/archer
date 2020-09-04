import * as ex from "excalibur"
import { Planet } from "./planet"
import * as res from "./resources"
import { getCookie } from "./utils"
import { ScreenSelector } from "./screen_selector"
import { Consts } from "./constants"
import { Indicator } from "./indicator"
import { circle } from "excalibur/dist/Util/DrawUtil"

export class Player extends ex.Actor {
    playerColor: ex.Color
    activated: boolean
    posChanged: boolean
    timer: ex.Timer
    ss: ScreenSelector
    username: string
    game: ex.Engine
    desiredAlpha: number
    planet: Planet
    ind: Indicator

    constructor(username: string, p: Planet, alpha: number, colorID: number, ss: ScreenSelector, game: ex.Engine) {
        super()

        this.desiredAlpha = -1
        this.game = game
        this.username = username
        this.ss = ss
        this.posChanged = false
        this.activated = false
        this.playerColor = res.Colors[colorID]
        var playerSprite = res.Images.player.asSprite().clone()
        playerSprite.scale = new ex.Vector(0.1, 0.1)
        playerSprite.offset = new ex.Vector(0, 25 + p.radius)
        playerSprite.colorize(this.playerColor)
        this.addDrawing(playerSprite)
        this.rotation = alpha
        this.timer = new ex.Timer({
            interval: Consts.movInterval,
            fcn: () => {
                this.updatePosition()
            },
            repeats: true,
        })
        this.planet = p
        this.planet.add(this)
        this.ind = new Indicator(this.playerColor, - this.planet.radius - Consts.indicatorOffset)
    }

    public onPostUpdate(engine: ex.Engine, delta: number) {
        if (this.desiredAlpha != -1) {
            // Find direction and diff of shorter arc
            // connecting two angles.
            var mult = 1
            var diffPos = this.desiredAlpha - this.rotation
            if (diffPos < 0)
                diffPos += 2 * Math.PI
            var diffNeg = this.rotation - this.desiredAlpha
            if (diffNeg < 0)
                diffNeg += 2 * Math.PI

            if (diffNeg < diffPos)
                mult = -1

            var diff = Math.min(diffPos, diffNeg)
            var change = delta * Consts.omega
            if (change >= diff) {
                change = diff
                this.desiredAlpha = -1
            }
            this.rotation += mult * change
        }
        if (!this.activated) {
            return
        }

        if (engine.input.keyboard.isHeld(ex.Input.Keys.Right)) {
            this.rotation += delta * Consts.omega
            this.posChanged = true
        }
        if (engine.input.keyboard.isHeld(ex.Input.Keys.Left)) {
            this.rotation -= delta * Consts.omega
            this.posChanged = true
        }
    }

    public activate() {
        this.activated = true
        this.add(this.ind)
        this.game.add(this.timer)
    }

    public deactivate() {
        this.remove(this.ind)
        this.activated = false
        this.game.remove(this.timer)
    }

    updatePosition() {
        if (this.posChanged) {
            this.remove(this.ind)

            var gameID = getCookie("game_id")
            var userID = getCookie("user_id")

            $.post("/game/move", { "game_id": gameID, "user_id": userID, "username": this.username, "new_alpha": this.rotation }, () => {
                this.posChanged = false
            }, "json").fail(() => {
                this.ss.restoreToWelcomeScreen(true)
            })
        }
    }

    public setDestination(alpha: number) {
        if (alpha != this.rotation)
            this.desiredAlpha = alpha
    }

    public onPostDraw(ctx: CanvasRenderingContext2D, delta: number) {
        if (Consts.enableDebug) {
            var extraOffsets: number[] = [3, 15, 30, 40]
            for (let eo of extraOffsets) {
                circle(ctx, 0, -this.planet.radius - eo, 10, ex.Color.White)
            }
        }
    }

    killPlayer() {
        console.log(`${this.username} is killed`)
        this.planet.remove(this)
    }
}