package main

import (
	"bufio"
	"fmt"
	"github.com/ingemar-fei/pokedexcli/internal/PokeAPI"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(...string) error
}

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("Pokedex > ")
		reader.Scan()
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		commandArgs := []string{}
		if len(words) > 1 {
			commandArgs = words[1:]
		}
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(commandArgs...)
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
		"explore": {
			name: "explore",
			description: "explore an area and find all pokemons there",
			callback: commandExplore,
		},
	}
}

func commandExplore(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Provide a locationArea name to explore")
	}
	locationArea := args[0]
	fmt.Printf("Exploring : %s ...\n", locationArea)
	pokemons, err := PokeAPI.ExploreArea(locationArea)
	if err != nil {
		return err
	}
	fmt.Println("Found pokemon:")
	for _,pokemon := range pokemons {
		fmt.Printf(" - %s\n",pokemon)
	}
	return nil
}

func commandMapBack(args ...string) error {
	listAreas, err := PokeAPI.GetPreviousLocationAreas()
	if err != nil {
		return err
	}
	for _, area := range listAreas {
		fmt.Println(area)
	}
	return nil
}

func commandMap(args ...string) error {
	listAreas, err := PokeAPI.GetNextLocationAreas()
	if err != nil {
		return err
	}
	for _, area := range listAreas {
		fmt.Println(area)
	}
	return nil
}

func commandExit(args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(args ...string) error {
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
