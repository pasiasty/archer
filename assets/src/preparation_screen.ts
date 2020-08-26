import { Screen } from "./screen"

export class PreparationScreen extends Screen {
    constructor() {
        super("preparation_screen")

        const label = document.createElement("label")
        label.innerText = "Preparation screen"
        label.className = 'label'

        this.disable()
    }
}
