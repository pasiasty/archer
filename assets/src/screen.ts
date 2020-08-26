import { getElement } from "./utils"
import { ScreenSelector } from "./screen_selector"

export class Screen {
    ui: HTMLElement
    ss: ScreenSelector
    name: string

    constructor(id: string, ss: ScreenSelector) {
        this.name = id
        this.ui = getElement(id)
        this.ss = ss
    }

    disable() {
        this.ui.hidden = true
        this.ui.setAttribute("disabled", "true")
    }

    enable() {
        this.ui.hidden = false
        this.ui.setAttribute("disabled", "false")
    }
}
