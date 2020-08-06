module github.com/bols-blue-org/ddns-r53

go 1.14

require (
	github.com/aws/aws-sdk-go v1.33.20
	github.com/mitchellh/go-homedir v1.1.0
	github.com/pion/webrtc/v2 v2.2.23 // indirect
	github.com/pion/webrtc/v3 v3.0.0-beta.1 // indirect
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.4.0
	github.com/urfave/cli/v2 v2.2.0
	local.packages/awsr53 v0.0.0-00010101000000-000000000000
	local.packages/globalip v0.0.0-00010101000000-000000000000
)

replace local.packages/awsr53 => ./awsr53

replace local.packages/globalip => ./globalip
