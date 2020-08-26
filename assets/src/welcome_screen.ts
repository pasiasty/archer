import { Screen } from "./screen"

export class WelcomeScreen extends Screen {
    constructor() {
        super("welcome_screen")

        const container = document.createElement('div')
        container.className = 'ui_container'

        const createGame = document.createElement('button')
        createGame.className = 'button'
        createGame.innerText = "Create game"

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