import { getElement } from "./utils"

export class Screen {
    ui: HTMLElement

    constructor(id: string) {
        this.ui = getElement(id)
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
