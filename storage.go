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

func Categorias() []string {
	url := "https://api.mercadolibre.com/sites/MLA/categories"
	resp, err := http.Get(url)
	if err != nil {
		panic("Explotur")
	}
	defer resp.Body.Close()
	var body []map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		panic("at the dis...")
	}

	results := body[0]
	fmt.Println(results)

	return nil
}
func Hijos(category string) []string {
	url := "https://api.mercadolibre.com/categories/" + category
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

	results := body["children_categories"].([]interface{})
	total := body["total_items_in_this_category"].(float64)
	// var res []interface{}
	res := make([]string, len(results))
	var totales float64
	for i := range results {
		resi := results[i].(map[string]interface{})
		res[i] = resi["id"].(string)
		totales += resi["total_items_in_this_category"].(float64)
	}
	// fmt.Println(res)
	fmt.Println(totales, total, total == totales)

	return res
}
func Contador(category string) int {
	// url := "https://api.mercadolibre.com/sites/MLA/search?category=MLA5726&offset=50&limit=50"
	url := "https://api.mercadolibre.com/sites/MLA/search?category=" + category
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
