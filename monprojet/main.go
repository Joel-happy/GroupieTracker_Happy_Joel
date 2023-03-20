package main 

import (
	"encoding/json"
	
	"io/ioutil"
)

func getCocktails() ([]Cocktail, error) {
    resp, err := http.Get("https://www.thecocktaildb.com/api/json/v1/1/search.php?s=margarita")
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var response CocktailResponse

    err = json.Unmarshal(body, &response)
    if err != nil {
        return nil, err
    }

    return response.Cocktails, nil
}
