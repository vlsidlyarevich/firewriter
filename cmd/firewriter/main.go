package main

import "github.com/vlsidlyarevich/firewriter/web/app"

func main() {
	var server = app.Server{}
	server.Init()
	server.InitHandlers()
	server.Start("web/static", "8000")
}
