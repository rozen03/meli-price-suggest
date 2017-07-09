package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//"github.com/nobonobo/unqlitego"
// "github.com/tpotlog/unqlitego"
// "github.com/tpotlog/unqlitego/collections"
//"github.com/nobonobo/unqlitego/collections"
//"collections"

// "fmt"
//"errors"
type Paging struct {
	total  int
	offset int
	limit  int
}
type Results struct {
	site_id           string                   `json:"site_id"`
	paging            map[string]interface{}   `json:"paging"`
	results           []map[string]interface{} `json:"results"`
	secondary_results []map[string]interface{} `json:"secondary_results"`
	related_results   []map[string]interface{} `json:"related_results"`
	sort              map[string]interface{}   `json:"sort"`
	available_sorts   []map[string]interface{} `json:"available_sorts"`
	filters           []map[string]interface{} `json:"filters"`
	available_filters []map[string]interface{} `json:"available_filters"`
}
type Storage struct {
}

func Contador(category string) int {
	url := "https://api.mercadolibre.com/sites/MLA/search?category=MLA5726&offset=50&limit=50"
	resp, err := http.Get(url)
	if err != nil {
		panic("Explotur")
	}

	defer resp.Body.Close()
	var body map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		panic("at the dis...")
	}

	results := body["results"].([]interface{})
	fmt.Println(results[0])
	return 0
}
