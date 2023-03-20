


func getDataFromAPI() ([]Cocktail, error) {
    url := "https://www.thecocktaildb.com/api/json/v1/1/search.php?s=margarita"
    response, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()

    var cocktails CocktailResponse
    err = json.NewDecoder(response.Body).Decode(&cocktails)
    if err != nil {
        return nil, err
    }

    return cocktails.Cocktails, nil
}
