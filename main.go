package main

import (
	"fmt"
	"pokedexcli/internal/pokeapi"
	"log"
)



func main() {

	pokeapiClient := pokeapi.NewClient()

	resp , err := pokeapiClient.ListLocationAreas()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)

	// startRepl()

}
