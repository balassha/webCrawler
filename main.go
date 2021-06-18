package main

import (
	"htmlparser/controllers"
	"htmlparser/infrastructure/httpServer"
	"htmlparser/infrastructure/router"
)

func main() {
	//Initialize routes
	router := router.SetupRouter()

	//Initialize HTTP Client
	controllers.Initialize()

	//Launch HTTP server
	server := httpServer.InitHttpServer(router)
	httpServer.LaunchServer(server)

	//Enable Graceful shutdown of Http server
	httpServer.ListenForInterruptsAndShutdownGracefully(server)
}
