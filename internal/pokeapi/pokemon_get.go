package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	if pokemonName == "" {
		fmt.Println("Missing Pokemon name!")
		return Pokemon{}, fmt.Errorf("missing parameter")
	}

	url := baseURL + "pokemon/" + pokemonName

	if val, ok := c.cache.Get(url); ok {
		pokemonResp := Pokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, fmt.Errorf("unsuccesful request")
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return Pokemon{}, fmt.Errorf("invalid pokemon name")
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonResp := Pokemon{}
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return Pokemon{}, nil
	}

	c.cache.Add(url, data)
	return pokemonResp, nil
}
