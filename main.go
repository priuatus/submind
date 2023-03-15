package main

import (
	"github.com/priuatus/submind/cmd"
	"github.com/priuatus/submind/config"
	"github.com/priuatus/submind/log"
	"github.com/samber/lo"
)

func main() {
	// prepare config and logs
	lo.Must0(config.Init())
	lo.Must0(log.Init())

	// run the app
	cmd.Execute()
}
