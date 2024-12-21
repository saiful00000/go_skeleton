package main

import (
	"go_clean/infrastructure/containers"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	containers.Serve(g)
}