package main

import (
	"fmt"
	"log"

	"github.com/gobuffalo/envy"
	"github.com/pasiasty/archer/actions"
)

// main is the starting point for your Buffalo application.
// You can feel free and add to this `main` method, change
// what it does, etc...
// All we ask is that, at some point, you make sure to
// call `app.Serve()`, unless you don't want to start your
// application that is. :)
func main() {
	fmt.Println("ARCHER_DATABASE", envy.Get("ARCHER_DATABASE", ""))
	fmt.Println("ARCHER_DATABASE_USER", envy.Get("ARCHER_DATABASE_USER", ""))
	fmt.Println("ARCHER_DATABASE_PASSWORD", envy.Get("ARCHER_DATABASE_PASSWORD", ""))
	fmt.Println("ARCHER_DATABASE_HOST", envy.Get("ARCHER_DATABASE_HOST", ""))
	fmt.Println("ARCHER_DATABASE_PORT", envy.Get("ARCHER_DATABASE_PORT", ""))
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}

/*
# Notes about `main.go`

## SSL Support

We recommend placing your application behind a proxy, such as
Apache or Nginx and letting them do the SSL heavy lifting
for you. https://gobuffalo.io/en/docs/proxy

## Buffalo Build

When `buffalo build` is run to compile your binary, this `main`
function will be at the heart of that binary. It is expected
that your `main` function will start your application using
the `app.Serve()` method.

*/
