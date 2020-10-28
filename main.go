package main

import (
	"fake-web-pasar-api/controller"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	routes := gin.Default()
	port := os.Getenv("PORT")

	v1 := routes.Group("/api/v1/")

	v1.POST("get_detail_pasar", controller.GetDetailPasar)
	v1.POST("get_all_pasar", controller.GetListPasar)
	v1.POST("get_all_product", controller.GetAllProduct)
	fmt.Printf("Listening on port localhost:%s", port)
	routes.Run(port)
}