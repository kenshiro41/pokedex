package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"poke_api/db"
	"poke_api/pokemon"
	"poke_api/pokemon/entity"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	db := db.Connect()
	defer db.Close()
	// db.CreateTable(&entity.Pokemon{})
	bytes, err := ioutil.ReadFile("./pokedex.json")
	if err != nil {
		log.Fatal(err)
	}
	var pokemons []entity.Pokemon
	if err := json.Unmarshal(bytes, &pokemons); err != nil {
		log.Fatal(err)
	}
	for _, p := range pokemons {
		db.Table("pokemon").Create(&p)
	}
	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/", ip)
	r.GET("/pokemon", pokemon.Fetch)
	r.Run(":6001")
}

func ip(c *gin.Context) {
	i := c.ClientIP()
	c.JSON(200, i)
}
