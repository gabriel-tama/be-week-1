package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"example.com/shopifyx/config"
)

func main() {

	confVars, configErr := config.Get()

	if configErr != nil {
		log.Fatal(configErr)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(fmt.Sprintf("%s:%s", confVars.Host, confVars.Port))
}
