package actions

import (
	"fmt"
	"net/http"
)

func (as *ActionSuite) Test_HomeHandlerAsBot() {
	req := as.HTML("/")
	res := req.Get()
	as.Equal(http.StatusForbidden, res.Code)
}

func (as *ActionSuite) Test_HomeHandlerSuccessful() {
	req := as.HTML("/")
	req.Headers["User-Agent"] = "Mozilla"
	res := req.Get()
	as.Equal(http.StatusOK, res.Code)
}

func (as *ActionSuite) Test_JoinHomeHandlerJoinNonExisting() {
	req := as.HTML("/abc")
	req.Headers["Cookie"] = "game_id=abc;user_id=def"
	req.Headers["User-Agent"] = "Mozilla"
	res := req.Get()
	as.Equal(http.StatusNotFound, res.Code)
}

func (as *ActionSuite) Test_JoinHomeHandlerJoinExisting() {
	u := gm.CreateGame()
	req := as.HTML(fmt.Sprintf("/%s", u.GameID))
	req.Headers["User-Agent"] = "Mozilla"
	res := req.Get()
	as.Equal(http.StatusTemporaryRedirect, res.Code)
}
