package main

// import "fmt" // Lib for text formatting
import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil) // just for getting rid of the warning

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"bxkii": "1234",
		"benji": "1234",
	}))

	// Test request
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "testu lulu",
		})
	})

	// Login endpoint
	// TODO: add mongo support
	authorized.GET("/secret", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"secret": "Du bist so a geile sau",
		})
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run("127.0.0.1:8080")
}
