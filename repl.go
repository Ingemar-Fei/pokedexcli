package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ingemar-fei/pokedexcli/internal/PokeAPI"
	_ "github.com/ingemar-fei/pokedexcli/internal/PokeAPI"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays next 20 location-areas of the world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 location-areas of the world",
			callback:    commandMapBack,
		},
	}
}

func commandMapBack() error {
	listAreas, err := PokeAPI.GetPreviousLocationAreas()
	if err != nil {
		return err
	}
	for _, area := range listAreas {
		fmt.Println(area)
	}
	return nil
}

func commandMap() error {
	listAreas, err := PokeAPI.GetNextLocationAreas()
	if err != nil {
		return err
	}
	for _, area := range listAreas {
		fmt.Println(area)
	}
	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
