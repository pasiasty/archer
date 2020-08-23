require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
require("@fortawesome/fontawesome-free/js/all.js");

import * as ex from "excalibur/dist/excalibur.min.js"

import { MainMenu } from "./main_menu"

// Create an instance of the engine.
const game = new ex.Engine({isDebug: true})

// Start the engine to begin the game.
game.add('mainmenu', new MainMenu())
game.goToScene('mainmenu')
game.start()
