package PokeAPI

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

const PlayerNum = 100

func InspectPokemon(pokemon Pokemon) error {
	fmt.Printf(
		`Name: %s,
Height: %d,
Weight: %d,
Stats:
	-hp: %d,
	-attack: %d,
	-defense: %d,
	-special-attack: %d,
	-special-defense: %d,
	-speed: %d
`,
		pokemon.Name, pokemon.Height, pokemon.Weight, pokemon.Stats[0].Base_stat, pokemon.Stats[1].Base_stat, pokemon.Stats[2].Base_stat, pokemon.Stats[3].Base_stat, pokemon.Stats[4].Base_stat, pokemon.Stats[5].Base_stat)
	fmt.Println("Types:")
	for _, type_ := range pokemon.Types {
		fmt.Println(" - ", type_.Type.Name)
	}
	return nil
}

func TryCatchPokmon(pokemonName string) (Pokemon, error) {
	pokemon, err := getPokemon(pokemonName)
	if err != nil {
		return Pokemon{}, err
	}
	chance := pokemon.Base_experience
	playerRand := (rand.Intn(PlayerNum) * rand.Intn(PlayerNum)) % PlayerNum
	fmt.Println("Your rolled:", playerRand)
	pokemonRand := (rand.Intn(chance) * rand.Intn(chance)) % chance
	fmt.Println("Pokemon rolled:", pokemonRand)
	if playerRand > pokemonRand {
		return pokemon, nil
	} else {
		return Pokemon{}, fmt.Errorf("error : failed to catch pokemon %s", pokemonName)
	}
}

func getPokemon(pokemonName string) (Pokemon, error) {
	url := pokeAPIbase + "pokemon/" + pokemonName
	resp, err := callPokeAPI(url)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error: %w", err)
	}
	if string(resp) == "Not Found" {
		return Pokemon{}, fmt.Errorf("error : pokemon %s not found", pokemonName)
	}
	var reqRes Pokemon
	err = json.Unmarshal(resp, &reqRes)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error : get API json  %w", err)
	}
	return reqRes, nil
}
