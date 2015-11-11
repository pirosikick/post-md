package main

import (
	"fmt"
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandPost,
}

var commandPost = cli.Command{
	Name:        "post",
	Usage:       "Posts markdown",
	Description: "Posts markdown to sevices",
	Action:      doPost,
}

func doPost(c *cli.Context) {
	filePath := c.Args().Get(0)
	fmt.Println(filePath)
}
