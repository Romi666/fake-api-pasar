package main

import (
	"fake-web-pasar-api/controller"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	routes := gin.Default()
	port := os.Getenv("PORT")

	v1 := routes.Group("/api/v1/")

	v1.POST("get_detail_pasar", controller.GetDetailPasar)
	v1.POST("get_all_pasar", controller.GetListPasar)
	routes.Run(port)
}