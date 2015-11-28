package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"io/ioutil"
	//	"net/http"
	//	"net/http/httputil"
)

type QiitaConfig struct {
	Token string
}

type Config struct {
	Qiita QiitaConfig
}

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
	configStr, _ := ioutil.ReadFile("./.postmdrc")

	var config Config
	json.Unmarshal(configStr, &config)

	fmt.Printf("%#v\n", config)
	//	url := "https://qiita.com/api/v2/authenticated_user/items"
	//	req, _ := http.NewRequest("GET", url, nil)
	//	req.Header.Set("Authorization", "Bearer "+token)
	//
	//	client := new(http.Client)
	//	resp, _ := client.Do(req)
	//
	//	dumpResp, _ := httputil.DumpResponse(resp, true)
	//	fmt.Printf("%s", dumpResp)
}
