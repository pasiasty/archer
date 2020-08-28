package actions

import (
	"context"
	"testing"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/httptest"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/suite"
	"github.com/pasiasty/archer/server"
)

type ActionSuite struct {
	*suite.Action
	c buffalo.Context
	u *server.User
	g *server.Game
}

func (as *ActionSuite) BeforeTest(suiteName, testName string) {
	envy.Set("GO_ENV", "test")
	as.c = &buffalo.DefaultContext{
		Context: context.Background(),
	}

	gm = server.CreateGameManager()
	as.u = gm.CreateGame()
	as.g, _ = gm.GetGame(as.c, as.u.GameID)
}

func (as *ActionSuite) ProperRequest(path string) *httptest.Request {
	req := as.HTML(path)
	req.Headers["User-Agent"] = "Mozilla"
	req.Headers["X-CSRF-Token"] = "test"
	return req
}

func (as *ActionSuite) Request(path string, data map[string]interface{}, expRes int) *httptest.Response {
	req := as.ProperRequest(path)
	res := req.Post(data)
	as.Equal(expRes, res.Code)
	return res
}

func Test_ActionSuite(t *testing.T) {
	action, err := suite.NewActionWithFixtures(App(), packr.New("Test_ActionSuite", "../fixtures"))
	if err != nil {
		t.Fatal(err)
	}

	as := &ActionSuite{
		Action: action,
	}
	suite.Run(t, as)
}
