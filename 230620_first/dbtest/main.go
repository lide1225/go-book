package main

import (
	"dbt/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")
	router.InitRouter(r)
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
