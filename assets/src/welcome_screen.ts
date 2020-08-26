import { Screen } from "./screen"
import { ScreenSelector } from "./screen_selector"

export class WelcomeScreen extends Screen {
    constructor(ss: ScreenSelector) {
        super("welcome_screen", ss)

        const container = document.createElement('div')
        container.className = 'ui_container'

        const createGame = document.createElement('button')
        createGame.className = 'button'
        createGame.innerText = "Create game"

        createGame.onclick = () => {
            $.post("/preparation/create_game", (data) => {
                var resp = data as string
                if (resp != null && resp == "OK") {
                    this.ss.setCurrentScreen("preparation_screen")
                }
            })
        }

        const joinGame = document.createElement('button')
        joinGame.className = 'button'
        joinGame.innerText = "Join game"

        const input = document.createElement('input')
        input.className = 'input'

        container.appendChild(createGame)
        container.appendChild(input)
        container.appendChild(joinGame)

        this.ui.appendChild(container)
    }
}