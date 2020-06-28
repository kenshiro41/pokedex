package pokemon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"poke_api/db"
	"poke_api/pokemon/entity"

	"github.com/gin-gonic/gin"
)

func Fetch(c *gin.Context) {
	db := db.Connect()
	defer db.Close()

	bytes, err := ioutil.ReadFile("../pokedex.json")
	if err != nil {
		log.Fatal(err)
	}
	var pokemons []entity.Pokemon
	if err := json.Unmarshal(bytes, &pokemons); err != nil {
		log.Fatal(err)
	}
	for _, p := range pokemons {
		fmt.Printf("%d : %s\n", p.ID, p.Name)
	}
	c.JSON(200, "pokemon")
}
