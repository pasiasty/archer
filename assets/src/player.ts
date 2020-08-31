import * as ex from "excalibur"
import { Planet } from "./planet"
import * as res from "./resources"
import { getCookie } from "./utils"
import { ScreenSelector } from "./screen_selector"

const omega = 0.0025
const movInterval = 5

export class Player extends ex.Actor {
    colorID: number
    activated: boolean
    posChanged: boolean
    timer: ex.Timer
    ss: ScreenSelector
    username: string
    game: ex.Engine

    constructor(username: string, p: Planet, alpha: number, colorID: number, ss: ScreenSelector, game: ex.Engine) {
        super()

        this.game = game
        this.username = username
        this.ss = ss
        this.posChanged = false
        this.activated = false
        this.colorID = colorID
        var playerSprite = res.Images.player.asSprite().clone()
        playerSprite.scale = new ex.Vector(0.1, 0.1)
        playerSprite.offset = new ex.Vector(0, 25 + p.radius)
        this.addDrawing(playerSprite)
        this.rotation = alpha
        this.timer = new ex.Timer({
            interval: movInterval,
            fcn: () => {
                this.updatePosition()
            },
            repeats: true,
        })
        p.add(this)
    }

    public update(engine: ex.Engine, delta: number) {
        if (!this.activated) {
            return
        }

        if (engine.input.keyboard.isHeld(ex.Input.Keys.Right)) {
            this.rotation += delta * omega
            this.posChanged = true
        }
        if (engine.input.keyboard.isHeld(ex.Input.Keys.Left)) {
            this.rotation -= delta * omega
            this.posChanged = true
        }
    }

    public activate() {
        this.activated = true
        this.game.add(this.timer)
    }

    updatePosition() {
        if (this.posChanged) {
            var gameID = getCookie("game_id")
            var userID = getCookie("user_id")

            $.post("/game/move", { "game_id": gameID, "user_id": userID, "username": this.username, "new_alpha": this.rotation }, () => {
                this.posChanged = false
            }, "json").fail(() => {
                this.ss.restoreToWelcomeScreen()
            })
        }
    }
}