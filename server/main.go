package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	fmt.Println("starting server at port ", port)

	r := gin.Default()
	r.GET("/coffee", func(c *gin.Context) {
		handleCoffee()

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	r.Run(":" + port)
	// router.Run(":3000") for a hard coded port

}

func handleCoffee() {
	resp, err := http.Get("https://coffee.alexflipnote.dev/random.json")

	if err != nil {
		fmt.Println("Something is wrong. Please try again")
		log.Panic("Something is worng. Please try again")
		os.Exit(1)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(b))
}
