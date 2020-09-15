package actions

import (
	"encoding/json"
	"net/http"

	"github.com/pasiasty/archer/server"
)

func (as *ActionSuite) Test_Game_GetWorld() {
	as.Request("/game/get_world", map[string]interface{}{"game_id": "abc"}, http.StatusNotFound)
	as.Request("/game/get_world", map[string]interface{}{"game_id": as.u.GameID}, http.StatusForbidden)

	as.g.Start(as.c, as.u.UserID)

	as.Request("/game/get_world", map[string]interface{}{"game_id": as.u.GameID}, http.StatusOK)
}

func (as *ActionSuite) Test_Game_PollTurn() {
	as.Request("/game/poll_turn", map[string]interface{}{"game_id": "abc"}, http.StatusNotFound)
	as.Request("/game/poll_turn", map[string]interface{}{"game_id": as.u.GameID}, http.StatusForbidden)

	as.g.Start(as.c, as.u.UserID)

	res := as.Request("/game/poll_turn", map[string]interface{}{"game_id": as.u.GameID}, http.StatusOK)
	pollTurn := &server.PollTurn{}
	err := json.Unmarshal(res.Body.Bytes(), pollTurn)
	as.Equal(nil, err)

	w, err := as.g.GetWorld(as.c)
	as.Equal(nil, err)

	as.Equal(w.GetPublicWorld().CurrentPlayer.Name, pollTurn.CurrentPlayer)
}

func (as *ActionSuite) Test_Game_Move() {
}

func (as *ActionSuite) Test_Game_Shoot() {
}

func (as *ActionSuite) Test_Game_GetTrajectory() {
}
