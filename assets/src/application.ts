require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
require("@fortawesome/fontawesome-free/js/all.js");

import { WelcomeScreen } from "./welcome_screen"
import { PreparationScreen } from "./preparation_screen"
import { GameScreen } from "./game_screen"
import { ScreenSelector } from "./screen_selector"
import { getCookie } from "./utils"

$(() => {
    var ss = new ScreenSelector()

    var ws = new WelcomeScreen(ss)
    var ps = new PreparationScreen(ss)
    var game = new GameScreen(ss)

    ss.addScreen(ws)
    ss.addScreen(ps)
    ss.addScreen(game)

    var gameID = getCookie("game_id")
    var userID = getCookie("user_id")
    var username = getCookie("username")
    var gameStarted = getCookie("game_started")

    ss.setCurrentScreen("game_screen")
    // if (gameID != "" && userID != "" && username != "") {
    //     if (gameStarted == "true") {
    //         ss.setCurrentScreen("game_screen")
    //     }
    //     else
    //         ss.setCurrentScreen("preparation_screen")
    // } else {
    //     ss.setCurrentScreen("welcome_screen")
    // }
});
