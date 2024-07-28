package handlers

import (
	"fmt"
	"net/http"

	"github.com/Nixa001/micromart-services-product/internal/model"
	"github.com/Nixa001/micromart-services-product/internal/service"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
	products, err := h.service.ListProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(products)
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var newProduct model.Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Logique pour créer un nouveau produit
	c.JSON(http.StatusCreated, gin.H{"message": "Product created", "product": newProduct})
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	// Logique pour récupérer un produit spécifique
	c.JSON(http.StatusOK, gin.H{"message": "Get product details", "id": id})
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var updatedProduct model.Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Logique pour mettre à jour un produit
	c.JSON(http.StatusOK, gin.H{"message": "Product updated", "id": id})
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	// Logique pour supprimer un produit
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted", "id": id})
}

func (h *ProductHandler) SearchProducts(c *gin.Context) {
	query := c.Query("q")
	// Logique pour rechercher des produits
	c.JSON(http.StatusOK, gin.H{"message": "Search products", "query": query})
}

func (h *ProductHandler) GetProductsByCategory(c *gin.Context) {
	category := c.Param("category")
	// Logique pour obtenir les produits par catégorie
	c.JSON(http.StatusOK, gin.H{"message": "Get products by category", "category": category})
}

func (h *ProductHandler) UpdateStock(c *gin.Context) {
	id := c.Param("id")
	var stockUpdate struct {
		Quantity int `json:"quantity"`
	}
	if err := c.ShouldBindJSON(&stockUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Logique pour mettre à jour le stock
	c.JSON(http.StatusOK, gin.H{"message": "Stock updated", "id": id, "quantity": stockUpdate.Quantity})
}
