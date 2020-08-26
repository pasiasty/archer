import { getDiv } from "./utils"

export class PreparationScreen {
    ui: HTMLDivElement

    constructor() {
        this.ui = getDiv("preparation_screen")

        const label = document.createElement("label")
        label.innerText = "Preparation screen"
        label.className = 'label'

        this.ui.appendChild(label)
        this.ui.hidden = true
    }
}
