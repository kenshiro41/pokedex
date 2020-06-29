package main

import (
	"poke_api/pokemon"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	r := gin.New()
	config := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://"
		},
		MaxAge: 12 * time.Hour,
	})
	r.Use(config)
	r.Use(gin.Recovery())
	r.Use(static.Serve("/image", static.LocalFile("./images", true)))
	r.GET("/", ip)
	r.GET("/pokemon", pokemon.Fetch)
	r.GET("/pokemon/:id", pokemon.FetchById)
	r.Run(":6001")
}

func ip(c *gin.Context) {
	i := c.ClientIP()
	c.JSON(200, i)
}
