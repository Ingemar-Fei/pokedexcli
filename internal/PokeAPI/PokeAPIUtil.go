package PokeAPI

import (
	"fmt"
	"github.com/ingemar-fei/pokedexcli/internal/PokeCache"
	"io"
	"net/http"
	"time"
)

var apiCache *PokeCache.Cache
var map_page int

const pokeAPIbase = "https://pokeapi.co/api/v2/"

func init() {
	apiCache = PokeCache.NewCache(time.Minute * 1)
}

func callPokeAPI(url string) ([]byte, error) {
	if cacheData, ok := apiCache.Get(url); ok {
		fmt.Println("------------------------------")
		fmt.Println("-  Getting from pokeCache...")
		fmt.Println("------------------------------")
		return cacheData, nil
	}
	fmt.Println("------------------------------")
	fmt.Println("-  Getting from pokeAPI...")
	fmt.Println("------------------------------")
	httpData, err := httpGet(url)
	if err != nil {
		return []byte{}, err
	}
	apiCache.Add(url, httpData)
	return httpData, nil
}
func httpGet(url string) ([]byte, error) {
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
