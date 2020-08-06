package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
	cli "github.com/urfave/cli/v2"
	"local.packages/awsr53"
	"local.packages/globalip"
)

func action(c *cli.Context) error {
	// Default name for policy, role policy.
	ProfileName := c.String("profile")
	zoneid := c.String("zone")
	name := c.String("name")
	os.Setenv("AWS_PROFILE", ProfileName)

	fmt.Println("profile:", ProfileName, " zoneid:", zoneid, " host name:", name)

	// load credentials from ~/.aws/credentials
	// and region from ~/.aws/config.
	// コネクションの初期化
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// APIクライアント初期化
	svc := route53.New(sess)

	result, err := awsr53.GetHostedZone(svc, zoneid)
	if err != nil {
		return err
	}
	domain := result.HostedZone.Name
	t := time.NewTicker(30 * time.Second) // 30秒おきに通知
	defer func() {
		t.Stop()
	}()
	lastIP := ""
	fullDomain := name + "." + *domain
	for {
		select {
		case <-t.C:
			log.Println("2sec interval")
			addr := globalip.GetIPaddr()

			if addr == "error" {
				fmt.Println("can't get global ip address.")
			} else if lastIP == addr {
				fmt.Println("address not change :", addr)
			} else {
				fmt.Println("address change new:", addr, " old:", lastIP)
				awsr53.UpdateDNS(svc, fullDomain, addr, zoneid)
				lastIP = addr
			}
		}
	}
}
