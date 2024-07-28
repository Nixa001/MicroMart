package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Nixa001/micromart/api-gateway/internal/proxy"
	"github.com/Nixa001/micromart/pkg/common"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// CORS middleware
	router.Use(cors.Default())

	// API routes
	apiGroup := router.Group("/api")
	{
		productProxy := proxy.NewServiceProxy("product", common.GetServicePort("product"))
		apiGroup.Any("/product/*path", productProxy.HandleRequest)

		userProxy := proxy.NewServiceProxy("user", common.GetServicePort("user"))
		apiGroup.Any("/user/*path", userProxy.HandleRequest)

		cartProxy := proxy.NewServiceProxy("cart", common.GetServicePort("cart"))
		apiGroup.Any("/cart/*path", cartProxy.HandleRequest)

		paymentProxy := proxy.NewServiceProxy("payment", common.GetServicePort("payment"))
		apiGroup.Any("/payment/*path", paymentProxy.HandleRequest)
	}

	// Health check route
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	})

	port := os.Getenv("API_GATEWAY_PORT")
	if port == "" {
		port = "8080" // Default port
	}

	log.Printf("Starting API Gateway on port %s\n", port)
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Failed to start API Gateway: %v", err)
	}
}
