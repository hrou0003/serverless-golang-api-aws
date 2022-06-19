package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hrou0003/serverless-golang-api-aws/router"
)

func main() {
	r := gin.Default()
	router.LoadRouter(r)

	if err := r.Run(); err != nil {
		log.Printf("error starting server %+v", err)
	}
}
