package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) getData(url string) ([]byte, error) {
	if data, ok := c.cache.Get(url); ok {
		return data, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}

func (c *Client) GetLocations(pageURL *string) (e RespLocations, n error) {
	url := BaseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	data, err := c.getData(url)
	if err != nil {
		return RespLocations{}, err
	}

	respLocations := RespLocations{}
	if err := json.Unmarshal(data, &respLocations); err != nil {
		return RespLocations{}, err
	}

	c.cache.Add(url, data)
	return respLocations, nil
}

func (c *Client) GetLocationAreaPokemon(name *string) (Location, error) {
	url := BaseURL + "/location-area/" + *name
	data, err := c.getData(url)
	if err != nil {
		return Location{}, err
	}

	respLocation := Location{}
	if err := json.Unmarshal(data, &respLocation); err != nil {
		return Location{}, err
	}

	_, ok := c.cache.Get(url)
	if !ok {
		c.cache.Add(url, data)
	}

	return respLocation, nil
}

func (c *Client) GetPokemon(name *string) (Pokemon, error) {
	url := BaseURL + "/pokemon/" + *name
	data, err := c.getData(url)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, err
	}

	_, ok := c.cache.Get(url)
	if !ok {
		c.cache.Add(url, data)
	}

	return pokemon, nil
}
