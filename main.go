package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
	"local.packages/awsr53"
	"local.packages/globalip"
)

func main() {
	addr := globalip.GetIPaddr()

	if addr == "error" {
		fmt.Println("can't get global ip address.")
		return
	}

	fmt.Println("addr:", addr)
	// Default name for policy, role policy.
	RoleName := "ated"

	// Override name if provided
	if len(os.Args) == 2 {
		RoleName = os.Args[1]
	}
	os.Setenv("AWS_PROFILE", RoleName)

	// Initialize a session that the SDK uses to
	// load credentials from ~/.aws/credentials
	// and region from ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	zoneid := "Z089767717VJLDBYN3DHD"

	svc := route53.New(sess)

	result, err := awsr53.GetHostedZone(svc, zoneid)

	if err != nil {
		return
	}

	fmt.Println(result)

	result2, err := awsr53.ListResourceRecordSets(svc, zoneid)
	if err != nil {
		return
	}
	fmt.Println(result2)

	awsr53.UpdateDNS(svc, "test2.adrone.in.", addr, zoneid)

}
