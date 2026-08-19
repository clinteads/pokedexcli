package main

import (
	"fmt"
	"os"
)

var commands = map[string]cliCommand{
	"exit": {
		name: "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	},
	"help": {
		name: "help",
		description: "Displays a help message",
		callback: commandHelp,
	},
}

func commandMap(cfg *config) error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
    if err != nil {
        return err
    }
    defer res.Body.Close()

    body, err := io.ReadAll(res.Body)
    if err != nil {
        return err
    }

	var joke Joke
	if err := json.Unmarshal(body, &joke); err != nil {
		return err
	}

    // TODO: unmarshal body into a struct
    // TODO: loop and print names

    return nil
}

func commandExit(cfg *config) error{
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	fmt.Errorf("Error not reached")
	return nil
}

func commandHelp(cfg *config) error{
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")

	return nil
}