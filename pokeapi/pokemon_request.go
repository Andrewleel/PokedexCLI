package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullURL := baseURL + endpoint
	data, ok := c.cache.Get(fullURL)
	if ok {
		// Unmarshals "data" into valid struct which is locationAreasResp
		fmt.Println("cache hit!")
		pokemon := Pokemon{}
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}

		return pokemon, nil
	}
	fmt.Println("cache miss")
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code: %v ", resp.StatusCode)
	}
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	// Unmarshals "data" into valid struct which is locationAreasResp
	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	// add this URL to cache
	c.cache.Add(fullURL, data)
	return pokemon, nil
}