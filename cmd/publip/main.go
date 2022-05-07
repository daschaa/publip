package main

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type IpAddress struct {
	Ip string
}

func main() {
	app := &cli.App{
		Name:      "publip",
		Usage:     "Outputs the public ip adress",
		UsageText: "publip <parameter>",
		Authors:   []*cli.Author{{"Joshua Weber", "josh@joshuaw.de"}},
		Action: func(c *cli.Context) error {
			resp, err := http.Get("https://api.ipify.org?format=json")
			if err != nil {
				return err
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}
			var ipAddr IpAddress
			err = json.Unmarshal(body, &ipAddr)
			if err != nil {
				return err
			}
			fmt.Println(ipAddr.Ip)
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
