package actions

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
)

// PreparationScreen default implementation.
func PreparationScreen(c buffalo.Context) error {
	gameID, err := c.Cookies().Get("game_id")
	if err != nil {
		return errors.New("failed to fetch game_id from cookie")
	}
	userID, err := c.Cookies().Get("user_id")
	if err != nil {
		return errors.New("failed to fetch user_id from cookie")
	}
	c.Set("user_id", userID)
	c.Set("joining_link_url", fmt.Sprintf("http://%s/preparation/join_game?game_id=%s", envy.Get("SERVER_URL", "127.0.0.1:3000"), gameID))
	return c.Render(http.StatusOK, r.HTML("preparation/screen.html"))
}

// PreparationCreateGame default implementation.
func PreparationCreateGame(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.JSON("OK"))
}

// PreparationJoinGame default implementation.
func PreparationJoinGame(c buffalo.Context) error {
	gameID := c.Param("game_id")
	userID := gm.joinGame(gameID)
	if userID == -1 {
		return fmt.Errorf("game: %s does not exist", gameID)
	}
	c.Cookies().Set("game_id", gameID, 30*24*time.Hour)
	c.Cookies().Set("user_id", fmt.Sprint(userID), 30*24*time.Hour)
	return c.Redirect(302, "/preparation/screen")
}

// PreparationUserReady default implementation.
func PreparationUserReady(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("preparation/user_ready.html"))
}

// PreparationPollGame default implementation.
func PreparationPollGame(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("preparation/poll_game.html"))
}
