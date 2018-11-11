package main

import (
	"github.com/vlsidlyarevich/firewriter/web/app"
	"os"
)

func main() {
	var server = app.Server{}
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8000"
	}
	server.Init()
	server.InitHandlers()
	server.Start("web/static", port)
}
