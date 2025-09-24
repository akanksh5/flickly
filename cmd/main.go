package cmd

import (
	"context"
	"log"
	"net/http"

	"github.com/akanksh5/flickly/internal/services"
	"github.com/akanksh5/flickly/internal/stores"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var store stores.Store
	var service services.Service

	r := gin.Default()

	r.POST("/shorten", func(c *gin.Context) {
		var req struct {
			URL string `json:"url"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		shortURL, err := service.GenerateShortenedURL(context.Background(), req.URL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
	})

	r.GET("/expand/:short_url", func(c *gin.Context) {
		shortURL := c.Param("short_url")

		longURL, err := service.GetLongURL(context.Background(), shortURL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"long_url": longURL})
	})

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
