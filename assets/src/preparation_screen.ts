import { Screen } from "./screen"
import { getCookie, copyToClipboard, isHost, deleteCookie } from "./utils"
import { ScreenSelector } from "./screen_selector"

export class PreparationScreen extends Screen {
    joinLink: HTMLLabelElement
    usernameLabel: HTMLLabelElement
    timerID: number | undefined

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

        const leaveGame = document.createElement('button')
        leaveGame.className = 'preparation button'
        leaveGame.innerText = "Leave game"

        leaveGame.onclick = () => {
            deleteCookie("game_id")
            deleteCookie("user_id")
            deleteCookie("username")
            deleteCookie("is_host")
            this.ss.setCurrentScreen("welcome_screen")
        }

        container.appendChild(this.usernameLabel)
        container.appendChild(this.joinLink)
        container.appendChild(copyJoinLink)
        container.appendChild(leaveGame)

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

        this.timerID = window.setInterval(this.refresh, 1000)
    }

    disable() {
        super.disable()
        window.clearTimeout(this.timerID)
    }

    refresh() {
        var gameID = getCookie("game_id")

        $.post("/preparation/list_users", { "game_id": gameID }, (data) => {
            console.log(data)
        }, "json").fail((data) => {
            alert("Couldn't list users")
        })
    }
}
