require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
require("@fortawesome/fontawesome-free/js/all.js");

import { WelcomeScreen } from "./welcome_screen"
import { PreparationScreen } from "./preparation_screen"
import { GameScreen } from "./game_screen"

$(() => {

    var ws = new WelcomeScreen()
    var ps = new PreparationScreen()
    var game = new GameScreen()

});
