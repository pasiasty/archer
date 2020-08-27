import { Screen } from "./screen"
import { getCookie, copyToClipboard, isHost, deleteCookie } from "./utils"
import { ScreenSelector } from "./screen_selector"
import { UsersList } from "./messages"

export class PreparationScreen extends Screen {
    joinLink: HTMLLabelElement
    usernameLabel: HTMLLabelElement
    timerID: number | undefined
    userList: HTMLDivElement
    startGame: HTMLButtonElement
    userReady: HTMLButtonElement
    container: HTMLDivElement
    copyJoinLink: HTMLButtonElement
    addPlayer: HTMLButtonElement
    removePlayer: HTMLButtonElement
    leaveGame: HTMLButtonElement

    constructor(ss: ScreenSelector) {
        super("preparation_screen", ss)

        this.container = document.createElement('div')
        this.container.className = 'ui_container'

        this.usernameLabel = document.createElement("label")
        this.usernameLabel.className = 'preparation label'

        this.joinLink = <HTMLLabelElement>document.createElement("label")
        this.joinLink.className = 'preparation label'

        this.copyJoinLink = document.createElement('button')
        this.copyJoinLink.className = 'preparation button'
        this.copyJoinLink.innerText = "Copy link to clipboard"

        this.copyJoinLink.onclick = () => {
            copyToClipboard(this.joinLink.innerText)
        }

        this.userList = document.createElement('div')
        this.userList.className = 'preparation user_list'

        this.addPlayer = document.createElement('button')
        this.addPlayer.className = 'preparation button'
        this.addPlayer.innerText = "Add player"

        this.addPlayer.onclick = () => {
            this.postOrGoBack(this, "preparation/add_player")
        }

        this.removePlayer = document.createElement('button')
        this.removePlayer.className = 'preparation button'
        this.removePlayer.innerText = "Remove player"

        this.removePlayer.onclick = () => {
            this.postOrGoBack(this, "preparation/remove_player", false)
        }

        this.userReady = document.createElement('button')
        this.userReady.className = 'preparation button'
        this.userReady.innerText = "Ready"

        this.userReady.onclick = () => {
            this.postOrGoBack(this, "preparation/user_ready")
            this.container.removeChild(this.addPlayer)
            this.container.removeChild(this.removePlayer)
            this.container.removeChild(this.userReady)
        }

        this.startGame = document.createElement('button')
        this.startGame.className = 'preparation button'
        this.startGame.innerText = "Start game"

        this.startGame.onclick = () => {
        }

        this.leaveGame = document.createElement('button')
        this.leaveGame.className = 'preparation button'
        this.leaveGame.innerText = "Leave game"

        this.leaveGame.onclick = () => {
            this.restoreToWelcomeScreen()
        }

        this.ui.appendChild(this.container)
        this.disable()
    }

    prepareTopDescription(self: PreparationScreen) {

    }

    restoreToWelcomeScreen() {
        deleteCookie("game_id")
        deleteCookie("user_id")
        deleteCookie("username")
        deleteCookie("is_host")
        this.ss.setCurrentScreen("welcome_screen")
    }

    enable() {
        super.enable()
        var gameID = getCookie("game_id")
        var username = getCookie("username")
        this.joinLink.innerText = `${window.location.href}${gameID}`
        this.usernameLabel.innerText = `User: ${username}`

        this.container.appendChild(this.usernameLabel)
        this.container.appendChild(this.joinLink)
        this.container.appendChild(this.copyJoinLink)
        this.container.appendChild(this.userList)
        this.container.appendChild(this.addPlayer)
        this.container.appendChild(this.removePlayer)

        if (isHost()) {
            this.container.appendChild(this.startGame)
            this.usernameLabel.innerText += ' (host)'
        } else {
            this.container.appendChild(this.userReady)
        }

        this.container.appendChild(this.leaveGame)

        this.refresh(this)
        this.timerID = window.setInterval(this.refresh, 1000, this)
    }

    disable() {
        super.disable()

        while (this.container.firstChild) {
            if (this.container.lastChild != null)
                this.container.removeChild(this.container.lastChild);
        }
        window.clearTimeout(this.timerID)
    }

    updateUsersList(data: UsersList) {
        var lines: string[] = []
        for (let user of data.Users) {
            var userLines: string[] = []
            for (var i = 0; i < user.Players.length; i++) {
                var indent = '0px'
                if (i != 0) {
                    indent = '40px'
                }

                var colorText = ""

                if (user.IsHost) {
                    colorText = "#0000ff"
                } else if (user.Ready) {
                    colorText = "#00ff00"
                } else {
                    colorText = "#ff0000"
                }
                userLines.push(`<span style="color: ${colorText}; margin-left: ${indent};">${user.Players[i]}</span><br>`)
            }
            lines.push(userLines.join(""))
        }

        lines.sort((a: string, b: string): number => {
            return a.localeCompare(b)
        })

        this.userList.innerHTML = lines.join("")
    }

    refresh(self: PreparationScreen) {
        self.postOrGoBack(self, "preparation/list_users")
    }

    postOrGoBack(self: PreparationScreen, path: string, backOnFail = true) {
        var gameID = getCookie("game_id")
        var userID = getCookie("user_id")

        $.post(path, { "game_id": gameID, "user_id": userID }, (data: UsersList) => {
            self.updateUsersList(data)
        }, "json").fail(() => {
            if (backOnFail)
                self.restoreToWelcomeScreen()
        })
    }
}
