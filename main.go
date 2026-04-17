package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input_text := scanner.Text()
		clean_input := cleanInput(input_text)
		fmt.Printf("Your command was: %v\n",clean_input[0])
	}
}