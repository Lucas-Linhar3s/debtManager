package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Lucas-Linhar3s/debtManager/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/swaggo/swag/example/basic/docs"
)

// @title debtManager API
// @version 1.0
// @description This is the barber.

// @contact.name Linhares

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @host localhost:3333
// @BasePath /v1
func main() {
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Host = "localhost:3333"
	// Carrega as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	fmt.Println(os.Getenv("API_URL"))
	fmt.Println(os.Getenv("API_KEY"))

	r := gin.Default()

	interfaces.Routers(r.Group("/v1"))
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found " + " : " + c.Request.URL.String()})
	})

	if err := r.Run(":3333"); err != nil {
		log.Fatal(err)
	} // listen and serve on 0.0.0.0:8080

}
