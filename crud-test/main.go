package main

import (
	"crud-test/route"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	route.InitializeRoutes(r)

	err := r.Run(":3000")
	if err != nil {
		log.Fatalln(err)
	}
}
