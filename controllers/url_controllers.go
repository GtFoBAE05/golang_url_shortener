package controllers

import (
	"go-url-shortener/models"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func CreateUrl(c *gin.Context) {
	var body models.UrlRequestBody

	err := c.ShouldBindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, urlErr := url.ParseRequestURI(body.LongUrl)
	if urlErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": urlErr.Error()})
		return
	}

	err = models.CreateShortenUrl(body.LongUrl, body.ShortUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func RedirectToShortUrl(c *gin.Context) {
	shortUrl := c.Param("short_url")
	url, err := models.GetURLByShortURL(shortUrl)
	if err != nil {
		c.JSON(404, gin.H{"error": "No Url Found"})
		return
	}

	err = models.UpdateAccessCount(shortUrl)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url.LongUrl)
}

func GetUrlStats(c *gin.Context) {
	shortUrl := c.Param("short_url")
	url, err := models.GetURLByShortURL(shortUrl)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": url})

}
