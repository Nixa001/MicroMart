package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Récupérer le token du header Authorization
		token := c.GetHeader("Authorization")

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		// TODO: Vérifier la validité du token
		// Cette partie dépendra de votre implémentation spécifique (JWT, OAuth, etc.)
		// Par exemple:
		// if !isValidToken(token) {
		//     c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		//     c.Abort()
		//     return
		// }

		// Si tout est ok, continuer
		c.Next()
	}
}
