package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Nixa001/micromart/pkg/common"
	"github.com/gin-gonic/gin"
)

func main() {
	serviceName := "payment"
	port := common.GetServicePort(serviceName)
	router := gin.Default()

	// Routes spécifiques au service...
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}

	log.Printf("Starting %s service on port %s\n", serviceName, port)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
}
