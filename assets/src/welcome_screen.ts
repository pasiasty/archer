import { getDiv } from "./utils"

export class WelcomeScreen {
    ui: HTMLDivElement

    constructor() {
        this.ui = getDiv("welcome_screen")

        const createGame = document.createElement('button')
    createGame.className = 'button'
    createGame.id = "create_game"
    createGame.innerText = "Create game"

    this.ui.appendChild(createGame)
    }
}