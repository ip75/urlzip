package rest

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ip75/urlzip/pkg/core"
)

func RaiseHanders(g *gin.Engine, uz core.IUrlZip) {
	g.POST("/*uri", func(c *gin.Context) {
		shortPath, err := uz.ComposeShortURL(c.Request.RequestURI)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}

		scheme := "http"
		if c.Request.TLS != nil {
			scheme = "https"
		}

		shortURL := fmt.Sprintf("%s://%s/%s", scheme, c.Request.Host, shortPath)
		c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
	})

	g.GET("/:short", func(c *gin.Context) {
		shortPath := strings.ToLower(c.Param("short"))
		if !uz.Validate(shortPath) {
			c.Status(http.StatusNotFound)
			return
		}

		orig, err := uz.GetOriginalURL(shortPath)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.Redirect(http.StatusMovedPermanently, orig)
	})

}
