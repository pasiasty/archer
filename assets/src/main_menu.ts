require("expose-loader?$!expose-loader?jQuery!jquery");

import * as ex from "excalibur"

class Button extends ex.ScreenElement {
    constructor(text: string, x: number, y: number) {
        super({
            x: x,
            y: y,
            width: 200,
            height: 65,
            color: ex.Color.DarkGray,
        })

        var label = new ex.Label({
            text: text,
            x: 10,
            y: 50,
            fontSize: 28,
        })

        this.add(label)
    }

    onInitialize() {
        this.on('pointerup', () => {
            $.post("preparation/create_game")
        })
    }
}

export class MainMenu extends ex.Scene {
    public onInitialize(engine: ex.Engine) {
        this.add(new Button("Create game", 150, 150))
        this.add(new Button("Join game", 150, 250))
    }
}
