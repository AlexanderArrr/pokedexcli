package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) LocationArea(locationName string) (RespLocationArea, error) {
	if locationName == "" {
		fmt.Println("Missing area name! How to use explore:")
		fmt.Println("explore <area-name>")
		return RespLocationArea{}, fmt.Errorf("missing parameter")
	}

	url := baseURL + "location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		locationArea := RespLocationArea{}
		err := json.Unmarshal(val, &locationArea)
		if err != nil {
			return RespLocationArea{}, err
		}
		return locationArea, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationArea{}, fmt.Errorf("unsuccesful request")
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return RespLocationArea{}, fmt.Errorf("invalid area name")
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationArea{}, err
	}

	pokemonResp := RespLocationArea{}
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return RespLocationArea{}, nil
	}

	c.cache.Add(url, data)
	return pokemonResp, nil
}
