package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ingrid/routesapi"
)

func main() {
	r := gin.Default()

	r.GET("/routes", routesapi.GetFastestRoutes)

	if err := r.Run(); err != nil {
		fmt.Println("error running http handler", err)
	}
}
