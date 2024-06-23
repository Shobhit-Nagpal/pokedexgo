package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getLocation(url string) (MapResponse, []byte, error) {
	data := MapResponse{}
	res, err := http.Get(fmt.Sprintf("%s/location-area", url))
	if err != nil {
		fmt.Println(err.Error())
		return data, []byte{}, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
		return data, []byte{}, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err.Error())
		return data, []byte{}, err
	}

	return data, body, nil
}

func getPokemonsInArea(url, area string) (LocationAreaResponse, []byte, error) {
	data := LocationAreaResponse{}
	endpoint := fmt.Sprintf("%s/location-area/%s", url, area)
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err.Error())
		return data, []byte{}, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
		return data, []byte{}, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err.Error())
		return data, []byte{}, err
	}

	return data, body, nil
}

func getPokemonInfo(url, pokemon string) (PokemonResponse, []byte, error) {
	data := PokemonResponse{}
	endpoint := fmt.Sprintf("%s/pokemon/%s", url, pokemon)
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err.Error())
		return data, []byte{}, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
    fmt.Println(err.Error())
		return data, []byte{}, err
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err.Error())
		return data, []byte{}, err
	}
	return data, body, nil
}

func printLocations(locations MapResponse) {
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
}

func printPokemons(area LocationAreaResponse) {
	fmt.Println("Found Pokemon:")
	for _, encounter := range area.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}
}
