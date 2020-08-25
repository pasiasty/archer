require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
require("@fortawesome/fontawesome-free/js/all.js");

import * as ex from "excalibur"

import { MainMenu } from "./main_menu"

const game = new ex.Engine({})

game.add('mainmenu', new MainMenu(game))
game.goToScene('mainmenu')
game.start()
