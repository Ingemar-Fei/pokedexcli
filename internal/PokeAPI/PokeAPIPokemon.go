package PokeAPI

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

const PlayerNum = 100

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
		return Pokemon{}, fmt.Errorf("failed to catch pokemon %s", pokemonName)
	}
}

func getPokemon(pokemonName string) (Pokemon, error) {
	url := pokeAPIbase + "pokemon/" + pokemonName
	resp, err := callPokeAPI(url)
	if err != nil {
		return Pokemon{}, err
	}
	var reqRes Pokemon
	err = json.Unmarshal(resp, &reqRes)
	if err != nil {
		return Pokemon{}, err
	}
	return reqRes, nil
}
