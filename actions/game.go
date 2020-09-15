package actions

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/pasiasty/archer/models"
	"github.com/pasiasty/archer/server"

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
	return c.Render(http.StatusOK, r.JSON(world.GetPollTurn()))
}

// GameMove default implementation.
func GameMove(c buffalo.Context) error {
	gameID := c.Param("game_id")
	userID := c.Param("user_id")
	username := c.Param("username")
	newAlpha := c.Param("new_alpha")

	w, err := gm.GetWorld(c, gameID, userID, username)
	if err != nil {
		return err
	}

	newAlphaFloat, err := strconv.ParseFloat(newAlpha, 32)
	if err != nil {
		return c.Error(http.StatusBadRequest, fmt.Errorf("couldn't convert string to float: %v", err))
	}
	w.MovePlayer(c, username, float32(newAlphaFloat))
	return c.Render(http.StatusOK, r.JSON(server.Empty{}))
}

// GameShoot default implementation.
func GameShoot(c buffalo.Context) error {
	gameID := c.Param("game_id")
	userID := c.Param("user_id")
	username := c.Param("username")
	newAlpha := c.Param("new_alpha")
	shotX := c.Param("shot_x")
	shotY := c.Param("shot_y")

	w, err := gm.GetWorld(c, gameID, userID, username)
	if err != nil {
		return err
	}

	newAlphaFloat, err := strconv.ParseFloat(newAlpha, 32)
	if err != nil {
		return c.Error(http.StatusBadRequest, fmt.Errorf("couldn't convert string to float: %v", err))
	}
	w.MovePlayer(c, username, float32(newAlphaFloat))

	shotXFloat, err := strconv.ParseFloat(shotX, 32)
	if err != nil {
		return c.Error(http.StatusBadRequest, fmt.Errorf("couldn't convert string to float: %v", err))
	}

	shotYFloat, err := strconv.ParseFloat(shotY, 32)
	if err != nil {
		return c.Error(http.StatusBadRequest, fmt.Errorf("couldn't convert string to float: %v", err))
	}

	t, err := w.Shoot(c, username, server.Vector{X: float32(shotXFloat), Y: float32(shotYFloat)})
	if err != nil {
		return err
	}

	go func() {
		shot := &models.PlayerShot{
			GameID:     gameID,
			PlayerName: username,
			Collision:  t.CollidedWith,
		}
		vErrors, err := models.DB.ValidateAndSave(shot)
		if err != nil {
			c.Logger().Errorf("failed to put shot information into db: %v : %v", err, vErrors)
		}
	}()

	return c.Render(http.StatusOK, r.JSON(t))
}

// GameGetTrajectory default implementation.
func GameGetTrajectory(c buffalo.Context) error {
	gameID := c.Param("game_id")
	userID := c.Param("user_id")
	username := c.Param("username")

	w, err := gm.GetWorld(c, gameID, userID, username)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, r.JSON(w.GetTrajectory()))
}
