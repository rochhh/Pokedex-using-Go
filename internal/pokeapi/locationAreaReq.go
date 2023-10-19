package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ( c *Client ) ListLocationAreas() (LocationAreasResp , error){
	
	endpoint := "/location-area"
	fullurl := baseURL + endpoint

	req , err := http.NewRequest("GET" , fullurl , nil)

	if err != nil{
		return LocationAreasResp{} , err
	}

	resp , err := c.httpClient.Do(req)

	if err != nil{
		return LocationAreasResp{} , err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399{
		return LocationAreasResp{} , fmt.Errorf("Bad status code , %v " , resp.StatusCode)
	}


	data , err := io.ReadAll(resp.Body)

	if err != nil {
		return LocationAreasResp{} , err
	}

	locationAreasResp := LocationAreasResp{}
	
	err = json.Unmarshal(data , &locationAreasResp)

	if err != nil{
		return LocationAreasResp{} , err
	}

	return locationAreasResp , nil

}