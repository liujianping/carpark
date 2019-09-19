package main

import (
	_ "github.com/liujianping/carpark/http"
	_ "github.com/liujianping/carpark/job"
	_ "github.com/liujianping/carpark/prepare"
	"github.com/x-mod/build"
	"github.com/x-mod/cmd"
)

func main() {
	cmd.Version(build.String())
	cmd.Execute()
}
