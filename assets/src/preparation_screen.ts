import { Screen } from "./screen"
import { getCookie, copyToClipboard, isHost, deleteCookie, setCookie } from "./utils"
import { ScreenSelector } from "./screen_selector"
import { UsersList, GameStatus } from "./messages"
import { Consts } from "./constants"

export class PreparationScreen extends Screen {
    joinLink: HTMLLabelElement
    usernameLabel: HTMLLabelElement
    refreshTimerID: number | undefined
    gameStartedTimerID: number | undefined
    userList: HTMLDivElement
    startGame: HTMLButtonElement
    userReady: HTMLButtonElement
    container: HTMLDivElement
    copyJoinLink: HTMLButtonElement
    addPlayer: HTMLButtonElement
    removePlayer: HTMLButtonElement
    leaveGame: HTMLButtonElement
    moveTimeoutLabel: HTMLLabelElement
    moveTimeoutInput: HTMLInputElement
    enabled: boolean
    lastShootTimeout: number
    loopedWorldLabel: HTMLLabelElement
    loopedWorldCheckbox: HTMLInputElement

    constructor(ss: ScreenSelector) {
        super("preparation_screen", ss)
        this.enabled = false

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
            copyToClipboard(this.getJoinLink())
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

        this.moveTimeoutLabel = document.createElement('label')
        this.moveTimeoutLabel.className = 'preparation label'
        this.moveTimeoutLabel.innerText = 'Move timeout'

        this.lastShootTimeout = 0
        this.moveTimeoutInput = document.createElement('input')
        this.moveTimeoutInput.className = 'preparation'
        this.moveTimeoutInput.type = 'number'
        this.moveTimeoutInput.min = '0'
        this.moveTimeoutInput.step = '1'
        this.moveTimeoutInput.innerText = '0'
        this.moveTimeoutInput.addEventListener('keyup', (ev: Event) => {
            var v = parseInt(this.moveTimeoutInput.value)
            if (v == null || this.moveTimeoutInput.value == '')
                v = 0
            if (v != this.lastShootTimeout) {
                this.lastShootTimeout = v
                this.postGameSettings(v, this.loopedWorldCheckbox.checked)
            }
        })

        this.loopedWorldLabel = document.createElement('label')
        this.loopedWorldLabel.className = 'preparation label'
        this.loopedWorldLabel.innerText = 'Looped world'

        this.loopedWorldCheckbox = document.createElement('input')
        this.loopedWorldCheckbox.type = 'checkbox'
        this.loopedWorldCheckbox.className = 'preparation'
        this.loopedWorldCheckbox.addEventListener('click', () => {
            this.postGameSettings(this.lastShootTimeout, this.loopedWorldCheckbox.checked)
        })

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
            var gameID = getCookie("game_id")
            var userID = getCookie("user_id")
            var self = this
            $.post("preparation/start_game", { "game_id": gameID, "user_id": userID }, (data: UsersList) => {
                setCookie("game_started", "true")
                self.ss.setCurrentScreen("game_screen")
            }, "json").fail(() => {
                self.ss.restoreToWelcomeScreen()
            })
        }

        this.leaveGame = document.createElement('button')
        this.leaveGame.className = 'preparation button'
        this.leaveGame.innerText = "Leave game"

        this.leaveGame.onclick = () => {
            this.ss.restoreToWelcomeScreen()
        }

        this.ui.appendChild(this.container)
        this.disable()
    }

    postGameSettings(shootTimeout: number, loopedWorld: boolean) {
        var gameID = getCookie("game_id")
        var userID = getCookie("user_id")
        var self = this
        $.post("preparation/game_settings", { "game_id": gameID, "user_id": userID, "shoot_timeout": shootTimeout, "looped_world": loopedWorld }).fail(() => {
            alert('failed to set game settings!')
            self.moveTimeoutInput.value = ''
        })
    }

    getJoinLink(): string {
        var gameID = getCookie("game_id")
        return `${window.location.href}${gameID}`
    }

    enable() {
        super.enable()
        this.enabled = true
        var gameID = getCookie("game_id")
        var username = getCookie("username")
        this.joinLink.innerText = `Game ID: ${gameID}`
        this.usernameLabel.innerText = `User: ${username}`

        this.container.appendChild(this.usernameLabel)
        this.container.appendChild(this.joinLink)
        this.container.appendChild(this.copyJoinLink)
        this.container.appendChild(this.userList)
        this.container.appendChild(this.addPlayer)
        this.container.appendChild(this.removePlayer)

        this.container.appendChild(this.moveTimeoutLabel)
        this.container.appendChild(this.moveTimeoutInput)
        this.container.appendChild(this.loopedWorldLabel)
        this.container.appendChild(this.loopedWorldCheckbox)

        if (!isHost()) {
            this.moveTimeoutInput.disabled = true
            this.loopedWorldCheckbox.disabled = true
        }

        if (isHost()) {
            this.container.appendChild(this.startGame)
            this.usernameLabel.innerText += ' (host)'
        } else {
            this.container.appendChild(this.userReady)
        }

        this.container.appendChild(this.leaveGame)

        this.refresh(this)
        this.refreshTimerID = window.setTimeout(this.refresh, Consts.userListRefreshTimeout, this)

        if (!isHost())
            this.gameStartedTimerID = window.setTimeout(this.pollGameStatus, Consts.gameStartedRefreshTimeout, this)
    }

    disable() {
        super.disable()
        this.enabled = false

        while (this.container.firstChild) {
            if (this.container.lastChild != null)
                this.container.removeChild(this.container.lastChild);
        }
        window.clearTimeout(this.refreshTimerID)
        window.clearTimeout(this.gameStartedTimerID)
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
        var gameID = getCookie("game_id")
        var userID = getCookie("user_id")

        $.post("preparation/list_users", { "game_id": gameID, "user_id": userID }, (data: UsersList) => {
            self.updateUsersList(data)
            if (self.enabled)
                self.refreshTimerID = window.setTimeout(self.refresh, Consts.userListRefreshTimeout, self)
        }, "json").fail(() => {
            self.ss.restoreToWelcomeScreen()
        })
    }

    pollGameStatus(self: PreparationScreen) {
        var gameID = getCookie("game_id")

        $.post("/preparation/game_status", { "game_id": gameID }, (data: GameStatus) => {
            var ws = data.WorldSettings
            self.moveTimeoutInput.value = ws.ShootTimeout.toString()
            self.loopedWorldCheckbox.checked = ws.LoopedWorld
            if (data.Started == true) {
                setCookie("game_started", "true")
                self.ss.setCurrentScreen("game_screen")
            }
            if (self.enabled)
                self.gameStartedTimerID = window.setTimeout(self.pollGameStatus, Consts.gameStartedRefreshTimeout, self)
        }, "json").fail(() => {
            self.ss.restoreToWelcomeScreen()
        })
    }

    postOrGoBack(self: PreparationScreen, path: string, backOnFail = true) {
        var gameID = getCookie("game_id")
        var userID = getCookie("user_id")

        $.post(path, { "game_id": gameID, "user_id": userID }, (data: UsersList) => {
            self.updateUsersList(data)
        }, "json").fail(() => {
            if (backOnFail)
                self.ss.restoreToWelcomeScreen()
        })
    }
}
