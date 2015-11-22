package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"net/http"
	"net/http/httputil"
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
	token := c.Args().Get(0)
	fmt.Println(token)

	url := "https://qiita.com/api/v2/authenticated_user/items"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	client := new(http.Client)
	resp, _ := client.Do(req)

	dumpResp, _ := httputil.DumpResponse(resp, true)
	fmt.Printf("%s", dumpResp)
}
