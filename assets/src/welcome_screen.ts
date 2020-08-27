import { Screen } from "./screen"
import { ScreenSelector } from "./screen_selector"
import { setCookie } from "./utils"

export class WelcomeScreen extends Screen {
    constructor(ss: ScreenSelector) {
        super("welcome_screen", ss)

        const container = document.createElement('div')
        container.className = 'ui_container'

        const createGame = document.createElement('button')
        createGame.className = 'welcome button'
        createGame.innerText = "Create game"

        createGame.onclick = () => {
            $.post("/preparation/create_game", (data) => {
                this.ss.setCurrentScreen("preparation_screen")
            }, "json").fail((data) => {
                alert("Couldn't create game")
            })
        }

        const input = document.createElement('input')
        input.className = 'welcome input'

        const joinGame = document.createElement('button')
        joinGame.className = 'welcome button'
        joinGame.innerText = "Join game"

        joinGame.onclick = () => {
            $.post("/preparation/join_game", { "game_id": input.value }, (data) => {
                this.ss.setCurrentScreen("preparation_screen")
            }, "json").fail((data) => {
                alert("Couldn't find game: " + input.value)
            })
        }

        container.appendChild(createGame)
        container.appendChild(input)
        container.appendChild(joinGame)

        this.ui.appendChild(container)

        this.disable()
    }
}