package PokeAPI

import (
	"encoding/json"
	"fmt"
)

func ExploreArea(locationArea string) ([]string, error) {
	url := pokeAPIbase + "location-area/" + locationArea
	resp, err := callPokeAPI(url)
	if err != nil {
		return []string{}, err
	}
	var reqRes LocationArea
	err = json.Unmarshal(resp, &reqRes)
	if err != nil {
		return []string{}, err
	}
	var pokeRes []string
	for _, pokemon := range reqRes.Pokemon_encounters {
		pokeRes = append(pokeRes, pokemon.Pokemon.Name)
	}
	return pokeRes, nil
}

func GetNextLocationAreas() ([]string, error) {
	return getLocationAreas(true)
}

func GetPreviousLocationAreas() ([]string, error) {
	return getLocationAreas(false)
}

func getLocationAreas(forward bool) ([]string, error) {
	if forward {
		map_page++
	} else {
		if map_page <= 1 {
			return []string{}, fmt.Errorf("you're on the first page")
		}
		map_page--
	}
	LocationAreaAPI := pokeAPIbase + "location-area"
	limit := 20
	offset := (map_page - 1) * limit
	url := fmt.Sprintf("%s?limit=%d&offset=%d", LocationAreaAPI, limit, offset)
	resp, err := callPokeAPI(url)
	if err != nil {
		return []string{}, err
	}
	var listResult NamedAPIResourceList
	err = json.Unmarshal(resp, &listResult)
	if err != nil {
		return []string{}, err
	}
	if listResult.Next == "" {
		map_page--
	}
	var listLoactionAreas []string
	for _, locationArea := range listResult.Results {
		listLoactionAreas = append(listLoactionAreas, locationArea.Name)
	}
	return listLoactionAreas, nil
}
