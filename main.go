package main

// import "fmt" // Lib for text formatting
import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupRouter() *gin.Engine {
	// Establish database connection with MongoDB
	uri := os.Getenv("MONGODB_URI") // Get DB uri from env file
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable.")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	coll := client.Database("sample_mflix").Collection("movies")

	// Start gin server
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

	r.GET("/randomStop", func(c *gin.Context) {
		title := "A Corner in Wheat"
		var result bson.M
		err = coll.FindOne(context.TODO(), bson.D{{Key: "title", Value: title}}).Decode(&result)
		if err == mongo.ErrNoDocuments {
			fmt.Printf("No document was found with the title %s\n", title)
			return
		}
		if err != nil {
			panic(err)
		}
		jsonData, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}

		jsonData2 := []byte(jsonData)
		c.Data(200, "application/json", jsonData2)
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := setupRouter()
	r.Run("127.0.0.1:8080")
}
