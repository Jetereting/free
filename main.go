package main

import (
	"os"

	"github.com/moqsien/free/pkgs/runner"
	"github.com/moqsien/free/pkgs/sites"
)

func main() {
	rn := runner.NewRunner()
	rn.RegisterSite(sites.NewMianfieFQ())

	storage_dir := os.Getenv("FREE_VPN_DIR")
	rn.Run(storage_dir)
}
