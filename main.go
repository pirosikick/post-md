package main

import (
	"github.com/codegangsta/cli"
	"os"
)

var Version string = "0.0.1"

func main() {
	newApp().Run(os.Args)
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "post-md"
	app.Usage = "Posts markdown text to some web services"
	app.Version = Version
	app.Author = "pirosikick"
	app.Email = "pirosikick@gmail.com"
	app.Commands = Commands
	return app
}
