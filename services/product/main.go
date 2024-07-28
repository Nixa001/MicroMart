package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	handler "github.com/Nixa001/micromart-services-product/internal/handlers"
	"github.com/Nixa001/micromart-services-product/internal/repository"
	"github.com/Nixa001/micromart-services-product/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// Établir la connexion à MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)

	// Ping la base de données pour vérifier la connexion
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	// Sélectionner la base de données et la collection
	db := client.Database("product_db")
	collection := db.Collection("products")

	// Initialiser le repository, le service et le handler
	repo := repository.NewProductRepository(collection)
	svc := service.NewProductService(repo)
	handler := handler.NewProductHandler(svc)

	// Configuration de Gin
	r := gin.Default()

	// Définition des routes
	productRoutes := r.Group("/products")
	{
		productRoutes.GET("", handler.ListProducts)
		productRoutes.POST("", handler.CreateProduct)
		productRoutes.GET("/:id", handler.GetProduct)
		productRoutes.PUT("/:id", handler.UpdateProduct)
		productRoutes.DELETE("/:id", handler.DeleteProduct)
		productRoutes.GET("/search", handler.SearchProducts)
		productRoutes.GET("/category/:category", handler.GetProductsByCategory)
		productRoutes.PATCH("/:id/stock", handler.UpdateStock)
	}

	// Démarrage du serveur
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
