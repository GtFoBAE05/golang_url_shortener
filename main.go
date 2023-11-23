package main

import (
	"go-url-shortener/controllers"
	"go-url-shortener/database"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.ConnectPostgres()

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db not connected")
	}

	err = database.Migrate(db)
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.POST("/shorten", controllers.CreateUrl)
	router.GET("/:short_url", controllers.RedirectToShortUrl)
	router.GET("/:short_url/stats", controllers.GetUrlStats)

	router.Run()
}
