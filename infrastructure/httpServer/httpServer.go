package httpServer

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func InitHttpServer(router *gin.Engine) *http.Server {
	//Ideally port should be passed as an environment variable. For simplicity its
	//hardcoded now.
	//port := os.Getenv("httpPort")
	port := ":8011"
	return &http.Server{
		Addr:    port,
		Handler: router,
	}
}

func LaunchServer(server *http.Server) {
	go func() {
		// service connections
		if err := server.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()
}

func ListenForInterruptsAndShutdownGracefully(server *http.Server) {
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Initiated...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
