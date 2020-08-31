import { Screen } from "./screen"
import { deleteCookie } from "./utils"

export class ScreenSelector {
    screens: Map<string, Screen>;
    currentScreen: string;

    constructor() {
        this.currentScreen = ""
        this.screens = new Map<string, Screen>()
    }

    addScreen(s: Screen) {
        this.screens.set(s.name, s)
    }

    setCurrentScreen(name: string) {
        if (this.currentScreen != "") {
            var s = this.screens.get(this.currentScreen)
            if (s == null) {
                throw new Error("Unable to find screen: " + this.currentScreen)
            }
            s.disable()
        }
        var s = this.screens.get(name)
        if (s == null) {
            throw new Error("Unable to find screen: " + name)
        }
        s.enable()
        this.currentScreen = name
    }

    restoreToWelcomeScreen() {
        deleteCookie("game_id")
        deleteCookie("user_id")
        deleteCookie("username")
        deleteCookie("is_host")
        deleteCookie("game_started")
        this.setCurrentScreen("welcome_screen")
    }
}