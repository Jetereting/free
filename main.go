package main

import (
	"github.com/moqsien/free/pkgs/runner"
	"github.com/moqsien/free/pkgs/sites"
)

func main() {
	rn := runner.NewRunner()
	rn.RegisterSite(sites.NewMianfieFQ())
	rn.Run()
}
