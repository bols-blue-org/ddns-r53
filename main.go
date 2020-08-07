package main

import (
	"fmt"
	"log"
	"os"

	cli "github.com/urfave/cli/v2"
)

func getHostname() string {
	ret, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}
	return ret
}

var flagDefine = []cli.Flag{
	&cli.StringFlag{
		Name:    "name",
		Aliases: []string{"n"},
		Value:   getHostname(),
		Usage:   "host name (ドメインの先頭)",
		EnvVars: []string{"HOSTNAME"},
	},
	&cli.StringFlag{
		Name:    "zone",
		Aliases: []string{"z"},
		Value:   "Z089767717VJLDBYN3DHD",
		Usage:   "aws route53 zone id",
		EnvVars: []string{"ZONE_ID"},
	},
	&cli.StringFlag{
		Name:    "profile",
		Aliases: []string{"p"},
		Value:   "ated",
		Usage:   "aws profile name. ~/aws/config and ~/aws/profile",
		EnvVars: []string{"AWS_PROFILE"},
	},
}

func main() {
	app := &cli.App{
		Flags: flagDefine,
	}
	app.Action = action

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
