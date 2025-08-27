package PokeAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var map_page int

type NamedAPIResourceList struct {
	Count    int                `json:"count"`
	Next     string             `json:"next"`
	Previous string             `json:"previous"`
	Results  []NamedAPIResource `json:"results"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationArea struct {
	Id                     int                   `json:"id"`
	Name                   string                `json:"name"`
	Game_index             int                   `json:"game_index"`
	Encounter_method_rates []EncounterMethodRate `json:"encounter_method_rates"`
}

type EncounterMethodRate struct {
	Encounter_method NamedAPIResource          `json:"encounter_method"`
	Version_details  []EncounterVersionDetails `json:"version_details"`
}

type EncounterVersionDetails struct {
	Rate    int              `json:"rate"`
	Version NamedAPIResource `json:"version"`
}

type PokemonEncounter struct {
	Pokemon         NamedAPIResource         `json:"pokemon"`
	Version_details []VersionEncounterDetail `json:"version_details"`
}

type VersionEncounterDetail struct {
	Version           NamedAPIResource `json:"version"`
	Max_chance        int              `json:"max_chance"`
	Encounter_details []Encounter      `json:"encounter_details"`
}

type Encounter struct {
	Min_level        int                `json:"min_level"`
	Max_level        int                `json:"max_level"`
	Condition_values []NamedAPIResource `json:"condition_values"`
	Chance           int                `json:"chance"`
	Method           NamedAPIResource   `json:"method"`
}

func CallPokeAPI(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	resString, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return resString, nil
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
			return []string{"you're on the first page"}, nil
		}
		map_page--
	}
	LocationAreaAPI := "https://pokeapi.co/api/v2/location-area"
	limit := 20
	offset := (map_page - 1) * limit
	url := fmt.Sprintf("%s?limit=%d&offset=%d", LocationAreaAPI, limit, offset)
	resp, err := CallPokeAPI(url)
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
