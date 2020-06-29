package pokemon

import (
	"fmt"
	"poke_api/db"
	"poke_api/pokemon/entity"

	"github.com/gin-gonic/gin"
)

func Fetch(c *gin.Context) {
	db := db.Connect()
	defer db.Close()
	var p []*entity.Pokemon
	pokemons := db.Table("pokemon").Find(&p)
	c.JSON(200, pokemons)
}
func FetchById(c *gin.Context) {
	db := db.Connect()
	defer db.Close()
	id := c.Param("id")
	fmt.Println(id)
	var p entity.Pokemon
	pokemon := db.Table("pokemon").Where("id = ?", id).First(&p)

	c.JSON(200, pokemon)
}
