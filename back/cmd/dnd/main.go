package main

import "github.com/var10000/DnDHitPoints/back/internal/service"

func main() {
	startServices()
}

func startServices() {
	app := service.InitServices()
	app.Start()
}
