package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()


	// Rota de teste
	r.GET("/ping", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Inicia o servidor na porta 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}