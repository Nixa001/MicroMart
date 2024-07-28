package main

import (
	"log"

	handler "github.com/Nixa001/micromart-gateway/handlers"
	"github.com/Nixa001/micromart-gateway/model"
	"github.com/gin-gonic/gin"
)

func main() {
	gateway := model.NewGateway()
	handler := handler.NewGatewayHandler(gateway)

	router := gin.Default()

	router.Any("/products/*path", handler.HandleProduct)
	router.Any("/cart/*path", handler.HandleCart)
	router.Any("/payment/*path", handler.HandlePayment)
	router.Any("/users/*path", handler.HandleUser)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
