package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"poke_api/db"
	"poke_api/pokemon/entity"
)

func main() {
	db := db.Connect()
	defer db.Close()
	db.CreateTable(&entity.Pokemon{})
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
} 
