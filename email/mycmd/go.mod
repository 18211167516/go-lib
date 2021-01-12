module github.com/18211167516/go-lib/email/mycmd

go 1.14

require (
	github.com/18211167516/go-lib/email/mycmd/cmd v0.0.0-20210111093206-ad1c0852a41e
	github.com/jordan-wright/email v0.0.0-20210109023952-943e75fe5223 // indirect
	github.com/spf13/cobra v1.1.1 // indirect
	github.com/spf13/viper v1.7.1 // indirect
)

replace github.com/18211167516/go-lib/email/mycmd/cmd v0.0.0-20210111093206-ad1c0852a41e => ./cmd
