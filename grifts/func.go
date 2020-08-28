package grifts

import (
	. "github.com/markbates/grift/grift"
)

var _ = Namespace("cov", func() {

	Desc("func", "Task Description")
	Add("func", func(c *Context) error {
		return runCoverage("func")
	})

})
