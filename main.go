package main

// import "fmt" // Lib for text formatting
import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil) // just for getting rid of the warning

	// Test request
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "testu lulu",
		})
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run("127.0.0.1:8080")
}
