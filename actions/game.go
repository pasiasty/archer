package actions

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gobuffalo/buffalo"
	"github.com/pasiasty/archer/server"
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

	return c.Render(http.StatusOK, r.JSON(world.GetPublicWorld()))
}

// GamePollTurn default implementation.
func GamePollTurn(c buffalo.Context) error {
	gameID := c.Param("game_id")

	game, err := gm.GetGame(c, gameID)
	if err != nil {
		return err
	}

	world, err := game.GetWorld(c)
	if err != nil {
		return err
	}

	pw := world.GetPublicWorld()

	return c.Render(http.StatusOK, r.JSON(server.PollTurn{
		CurrentPlayer:      pw.CurrentPlayer.Name,
		CurrentPlayerAlpha: pw.CurrentPlayer.Alpha,
		ShotPerformed:      false,
	}))
}

// GameMove default implementation.
func GameMove(c buffalo.Context) error {
	gameID := c.Param("game_id")
	userID := c.Param("user_id")
	username := c.Param("username")
	newAlpha := c.Param("new_alpha")
	// shotX := c.Param("shot_x")
	// shotY := c.Param("shot_y")

	game, err := gm.GetGame(c, gameID)
	if err != nil {
		return err
	}

	user, err := game.GetUser(c, userID)
	if err != nil {
		return err
	}

	if !user.HasPlayer(username) {
		return c.Error(http.StatusBadRequest, fmt.Errorf("player: %s does not belong to user: %s", username, userID))
	}

	w, err := game.GetWorld(c)
	if err != nil {
		return err
	}

	newAlphaFloat, err := strconv.ParseFloat(newAlpha, 32)
	if err != nil {
		c.Error(http.StatusBadRequest, fmt.Errorf("couldn't convert string to float: %v", err))
	}
	w.MovePlayer(c, username, float32(newAlphaFloat))
	return c.Render(http.StatusOK, r.JSON(server.Empty{}))
}
