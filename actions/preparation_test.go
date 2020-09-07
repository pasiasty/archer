package actions

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pasiasty/archer/server"
)

var (
	hexRegexp = `[0-9a-f]+`
)

func (as *ActionSuite) Test_Preparation_CreateGame() {
	res := as.Request("/preparation/create_game", nil, http.StatusOK)
	setCookies := res.Header().Values("Set-Cookie")
	as.Equal(4, len(setCookies))

	gameIDRes := setCookies[0]
	userIDRes := setCookies[1]
	usernameRes := setCookies[2]
	isHostRes := setCookies[3]

	as.Regexp(fmt.Sprintf(`game_id=%s;`, hexRegexp), gameIDRes)
	as.Regexp(fmt.Sprintf(`user_id=%s;`, hexRegexp), userIDRes)
	as.Regexp(`username="[a-zA-Z ]+";`, usernameRes)
	as.Regexp(`is_host=true;`, isHostRes)
}

func (as *ActionSuite) Test_Preparation_JoinGame() {
	as.Request("/preparation/join_game", map[string]interface{}{"game_id": "abc"}, http.StatusNotFound)
	res := as.Request("/preparation/join_game", map[string]interface{}{"game_id": as.u.GameID}, http.StatusOK)

	setCookies := res.Header().Values("Set-Cookie")
	as.Equal(4, len(setCookies))

	gameIDRes := setCookies[0]
	userIDRes := setCookies[1]
	usernameRes := setCookies[2]
	isHostRes := setCookies[3]

	as.Regexp(fmt.Sprintf(`game_id=%s;`, as.u.GameID), gameIDRes)
	as.Regexp(fmt.Sprintf(`user_id=%s;`, hexRegexp), userIDRes)
	as.Regexp(`username="[a-zA-Z ]+";`, usernameRes)
	as.Regexp(`is_host=false;`, isHostRes)
}

func (as *ActionSuite) Test_Preparation_UserReady() {
	as.Request("/preparation/user_ready", map[string]interface{}{"game_id": "abc"}, http.StatusNotFound)
	as.Request("/preparation/user_ready", map[string]interface{}{"game_id": as.u.GameID, "user_id": "abc"}, http.StatusNotFound)
	as.Request("/preparation/user_ready", map[string]interface{}{"game_id": as.u.GameID, "user_id": as.u.UserID}, http.StatusOK)
}

func (as *ActionSuite) Test_Preparation_ListUsers() {
	gm.JoinGame(as.c, as.u.GameID)
	gm.JoinGame(as.c, as.u.GameID)

	as.Request("/preparation/list_users", map[string]interface{}{"game_id": "abc"}, http.StatusNotFound)
	res := as.Request("/preparation/list_users", map[string]interface{}{"game_id": as.u.GameID}, http.StatusOK)
	usersList := &server.UsersList{}
	err := json.Unmarshal(res.ResponseRecorder.Body.Bytes(), usersList)
	as.Equal(err, nil)
	as.Equal(3, len(usersList.Users))
}

func (as *ActionSuite) Test_Preparation_AddPlayer() {
	as.Request("/preparation/add_player", map[string]interface{}{"game_id": "abc"}, http.StatusNotFound)
	as.Request("/preparation/add_player", map[string]interface{}{"game_id": as.u.GameID, "user_id": "abc"}, http.StatusNotFound)

	as.Equal(1, len(as.g.GetUsersList()[0].Players))

	as.Request("/preparation/add_player", map[string]interface{}{"game_id": as.u.GameID, "user_id": as.u.UserID}, http.StatusOK)

	as.Equal(2, len(as.g.GetUsersList()[0].Players))

	as.Request("/preparation/add_player", map[string]interface{}{"game_id": as.u.GameID, "user_id": as.u.UserID}, http.StatusOK)

	as.Equal(3, len(as.g.GetUsersList()[0].Players))

	as.u.MarkReady()

	as.Request("/preparation/add_player", map[string]interface{}{"game_id": as.u.GameID, "user_id": as.u.UserID}, http.StatusForbidden)
}

func (as *ActionSuite) Test_Preparation_RemovePlayer() {
	as.Request("/preparation/remove_player", map[string]interface{}{"game_id": "abc"}, http.StatusNotFound)
	as.Request("/preparation/remove_player", map[string]interface{}{"game_id": as.u.GameID, "user_id": "abc"}, http.StatusNotFound)
	as.Request("/preparation/remove_player", map[string]interface{}{"game_id": as.u.GameID, "user_id": as.u.UserID}, http.StatusForbidden)
	as.Request("/preparation/add_player", map[string]interface{}{"game_id": as.u.GameID, "user_id": as.u.UserID}, http.StatusOK)
	as.Request("/preparation/add_player", map[string]interface{}{"game_id": as.u.GameID, "user_id": as.u.UserID}, http.StatusOK)
	as.Request("/preparation/remove_player", map[string]interface{}{"game_id": as.u.GameID, "user_id": as.u.UserID}, http.StatusOK)

	res := as.Request("/preparation/list_users", map[string]interface{}{"game_id": as.u.GameID}, http.StatusOK)
	usersList := &server.UsersList{}
	err := json.Unmarshal(res.ResponseRecorder.Body.Bytes(), usersList)
	as.Equal(nil, err)
	as.Equal(2, len(usersList.Users[0].Players))

	as.u.MarkReady()

	as.Request("/preparation/remove_player", map[string]interface{}{"game_id": as.u.GameID, "user_id": as.u.UserID}, http.StatusForbidden)
}

func (as *ActionSuite) Test_Preparation_StartGame() {
	as.Equal(false, as.g.Status().Started)

	as.Request("/preparation/start_game", map[string]interface{}{"game_id": "abc"}, http.StatusNotFound)
	as.Request("/preparation/start_game", map[string]interface{}{"game_id": as.u.GameID, "user_id": "abc"}, http.StatusForbidden)
	as.Request("/preparation/start_game", map[string]interface{}{"game_id": as.u.GameID, "user_id": as.u.UserID}, http.StatusOK)

	as.Equal(true, as.g.Status().Started)
}

func (as *ActionSuite) Test_Preparation_GameStatus() {
	as.Request("/preparation/game_status", map[string]interface{}{"game_id": "abc"}, http.StatusNotFound)
	res := as.Request("/preparation/game_status", map[string]interface{}{"game_id": as.u.GameID}, http.StatusOK)
	var started server.GameStatus
	err := json.Unmarshal(res.ResponseRecorder.Body.Bytes(), &started)
	as.Equal(nil, err)
	as.Equal(server.GameStatus{Started: false}, started)

	as.g.Start(as.c, as.u.UserID)

	res = as.Request("/preparation/game_status", map[string]interface{}{"game_id": as.u.GameID}, http.StatusOK)
	err = json.Unmarshal(res.ResponseRecorder.Body.Bytes(), &started)
	as.Equal(nil, err)
	as.Equal(server.GameStatus{Started: true}, started)
}

func (as *ActionSuite) Test_Preparation_GameSettings() {
}
