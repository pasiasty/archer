import { Screen } from "./screen"
import { getCookie, copyToClipboard } from "./utils"
import { ScreenSelector } from "./screen_selector"
import { join } from "core-js/fn/array"

export class PreparationScreen extends Screen {
    joinLink: HTMLLabelElement

    constructor(ss: ScreenSelector) {
        super("preparation_screen", ss)

        const container = document.createElement('div')
        container.className = 'ui_container'

        const label = document.createElement("label")
        label.innerText = "Preparation screen"
        label.className = 'preparation label'

        this.joinLink = <HTMLLabelElement>document.createElement("label")
        this.joinLink.className = 'preparation label'

        const copyJoinLink = document.createElement('button')
        copyJoinLink.className = 'preparation button'
        copyJoinLink.innerText = "Copy link to clipboard"

        copyJoinLink.onclick = () => {
            copyToClipboard(this.joinLink.innerText)
        }

        container.appendChild(label)
        container.appendChild(this.joinLink)
        container.appendChild(copyJoinLink)

        this.ui.appendChild(container)

        this.disable()
    }

    enable() {
        super.enable()
        var gameID = getCookie("game_id")
        this.joinLink.innerText = `${window.location.href}${gameID}`
    }
}
