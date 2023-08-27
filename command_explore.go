package main

import (
	"fmt"
	"errors"
)


func callbackExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided ")
	}

	locationAreaName := args[0]
	// pokeapiClient := pokeapi.NewClient()
	// resp, err := pokeapiClient.ListLocationAreas()
	locationArea, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)

	if err != nil {
		return err
	}

	fmt.Printf("Pokemon in %v: \n", locationArea.Name)
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return nil
}