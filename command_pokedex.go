package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {

	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("You have not cuaght any Pokemon yet.")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for pokemonName := range cfg.caughtPokemon {
		fmt.Println(" - " + pokemonName)
	}


	return nil
}
