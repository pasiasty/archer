package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// GameGetWorld default implementation.
func GameGetWorld(c buffalo.Context) error {
	gameID := c.Param("game_id")

	game, err := gm.GetGame(c, gameID)
	if err != nil {
		return err
	}

	world, err := game.GetWorld(c)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, r.JSON(world))
}
