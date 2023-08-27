package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}
	startTime := time.Now()
	data, ok := c.cache.Get(fullURL)
	if ok {
		// Unmarshals "data" into valid struct which is locationAreasResp
		fmt.Println("cache hit!")
		locationAreasResp := LocationAreasResp{}
		err := json.Unmarshal(data, &locationAreasResp)
		endTime := time.Since(startTime)
		fmt.Printf("Request with Cache took %v \n", endTime)
		if err != nil {
			return LocationAreasResp{}, err
		}

		return locationAreasResp, nil
	}
	fmt.Println("cache miss")
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}

	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %v ", resp.StatusCode)
	}
	startTime = time.Now()
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	// Unmarshals "data" into valid struct which is locationAreasResp
	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}
	endTime := time.Since(startTime)
	fmt.Printf("Request without Cache took %v\nmap", endTime)

	// add this URL to cache
	c.cache.Add(fullURL, data)
	return locationAreasResp, nil
}


func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint
	data, ok := c.cache.Get(fullURL)
	if ok {
		// Unmarshals "data" into valid struct which is locationAreasResp
		fmt.Println("cache hit!")
		LocationAreaTemp := LocationArea{}
		err := json.Unmarshal(data, &LocationAreaTemp)
		if err != nil {
			return LocationArea{}, err
		}

		return LocationAreaTemp, nil
	}
	fmt.Println("cache miss")
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v ", resp.StatusCode)
	}
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	// Unmarshals "data" into valid struct which is locationAreasResp
	LocationAreaTemp := LocationArea{}
	err = json.Unmarshal(data, &LocationAreaTemp)
	if err != nil {
		return LocationArea{}, err
	}
	// add this URL to cache
	c.cache.Add(fullURL, data)
	return LocationAreaTemp, nil
}