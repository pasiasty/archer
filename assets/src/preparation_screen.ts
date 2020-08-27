import { Screen } from "./screen"
import { getCookie, copyToClipboard, isHost, deleteCookie } from "./utils"
import { ScreenSelector } from "./screen_selector"
import { UsersList } from "./messages"

export class PreparationScreen extends Screen {
    joinLink: HTMLLabelElement
    usernameLabel: HTMLLabelElement
    timerID: number | undefined
    userList: HTMLDivElement

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

        this.userList = document.createElement('div')
        this.userList.className = 'preparation user_list'

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
        container.appendChild(this.userList)
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

        this.refresh(this)
        this.timerID = window.setInterval(this.refresh, 1000, this)
    }

    disable() {
        super.disable()
        window.clearTimeout(this.timerID)
    }

    updateUsersList(data: UsersList) {
        var lines: string[] = []
        for (let user of data.Users) {
            var beginning = '<span style="color: '
            var postStyle = '">'
            var ending = '</span><br>'
            var colorText = ""

            if (user.IsHost) {
                colorText = "#0000ff"
            } else if (user.Ready) {
                colorText = "#00ff00"
            } else {
                colorText = "#ff0000"
            }
            lines.push(`${beginning}${colorText}${postStyle}${user.Username}${ending}`)
        }

        lines.sort((a: string, b: string): number => {
            return a.localeCompare(b)
        })

        this.userList.innerHTML = lines.join()
    }

    refresh(self: PreparationScreen) {
        var gameID = getCookie("game_id")

        $.post("/preparation/list_users", { "game_id": gameID }, (data: UsersList) => {
            console.log(typeof data, data)
            self.updateUsersList(data)
        }, "json").fail(() => {
            alert("Couldn't list users")
        })
    }
}
