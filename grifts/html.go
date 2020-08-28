package grifts

import (
	. "github.com/markbates/grift/grift"
)

var _ = Namespace("cov", func() {

	Desc("html", "Task Description")
	Add("html", func(c *Context) error {
		return runCoverage("html")
	})

})
