import { Screen } from "./screen"
import { getCookie, copyToClipboard, isHost } from "./utils"
import { ScreenSelector } from "./screen_selector"

export class PreparationScreen extends Screen {
    joinLink: HTMLLabelElement
    usernameLabel: HTMLLabelElement

    constructor(ss: ScreenSelector) {
        super("preparation_screen", ss)

        const container = document.createElement('div')
        container.className = 'ui_container'

        this.usernameLabel = document.createElement("label")
        this.usernameLabel.className = 'preparation label'

        this.joinLink = <HTMLLabelElement>document.createElement("label")
        this.joinLink.className = 'preparation label'

        const copyJoinLink = document.createElement('button')
        copyJoinLink.className = 'preparation button'
        copyJoinLink.innerText = "Copy link to clipboard"

        copyJoinLink.onclick = () => {
            copyToClipboard(this.joinLink.innerText)
        }

        container.appendChild(this.usernameLabel)
        container.appendChild(this.joinLink)
        container.appendChild(copyJoinLink)

        this.ui.appendChild(container)

        this.disable()
    }

    enable() {
        super.enable()
        var gameID = getCookie("game_id")
        var username = getCookie("username")
        this.joinLink.innerText = `${window.location.href}${gameID}`
        this.usernameLabel.innerText = `User: ${username}`

        if (isHost()) {
            this.usernameLabel.innerText += ' (host)'
        }
    }
}
