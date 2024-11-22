pa middlewares

import (
	"github.com/gin-gonic/gin"
	"notes_project/utils"
	"net/http"
	"strings"
)

// Middleware untuk memvalidasi token JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		// Memastikan token dimulai dengan "Bearer "
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Menghapus prefix "Bearer "
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		userID, err := utils.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Menyimpan userID ke context untuk digunakan di handler
		c.Set("userID", userID)
		c.Next()
	}
}
