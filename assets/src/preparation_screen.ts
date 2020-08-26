import { Screen } from "./screen"
import { ScreenSelector } from "./screen_selector"

export class PreparationScreen extends Screen {
    constructor(ss: ScreenSelector) {
        super("preparation_screen", ss)

        const label = document.createElement("label")
        label.innerText = "Preparation screen"
        label.className = 'label'

        this.ui.appendChild(label)

        this.disable()
    }
}
