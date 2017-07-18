package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
)

type obtainedData struct {
	min   float64
	max   float64
	sum   float64
	total float64
}

const melink = "https://api.mercadolibre.com/sites/MLA/search?limit=200&category="

/*
**Download from MeLi with the arguments recieved.
**Returns map[string]interface{} parsing with the response from MeLi.
 */
func Download(args string) map[string]interface{} {

	var resp *http.Response
	var err error
	var body map[string]interface{}
	//Download the json object from MeLi
	//Retry 100 times if error
	failed := 0
	for failed < 100 {
		resp, err = http.Get(melink + args)
		if err != nil {
			failed++
			continue
		}

		err = json.NewDecoder(resp.Body).Decode(&body)
		if err != nil {
			failed++
			continue
		}
		//Decode response ioreader into map[string]interface{}
		//decode to any struct would likely cause errors
		defer resp.Body.Close()
	}
	if err != nil {
		panic(err)
	}
	return body
}

/*
**Get the total count, sum, min and max of all prices.
 */
func GetPricesAndSold(items []interface{}) obtainedData {
	//Initialize return data
	prices := 0.0
	total := 0.0
	max := 0.0
	min := math.MaxFloat64
	/*
	**Loop over all items recieved
	**Sum the total and the prices
	**and evaluate the min and maximum values
	 */
	for i := range items {
		item := items[i].(map[string]interface{})
		price, ok := item["price"].(float64)
		if !ok {
			continue
		}
		sold, ok := item["sold_quantity"].(float64)
		if !ok {
			continue
		}
		max = math.Max(price, max)
		min = math.Min(price, min)
		prices += price * (sold + 1)
		total += (sold + 1)
	}
	return obtainedData{min, max, prices, total}
}

/*
**Download and returns min, max, total and prices
**From the recieved arguments
 */
func GetObtainedData(args string, c chan obtainedData, download Downloader) {
	body := download(args)
	results, ok := body["results"].([]interface{})
	if !ok {
		c <- obtainedData{0.0, 0.0, 0.0, 0.0}
		fmt.Println("Dio Feito :O")
	}
	go func() { c <- GetPricesAndSold(results) }()
}

/*
**Auxiliary method created for aesthetic code purposes
 */
func GetTotalCount(body *map[string]interface{}) int {
	return int((*body)["paging"].(map[string]interface{})["total"].(float64))
}

/*
**Send to the channel the arguments to download and process the items
**and also the channel where the task worker should send the prices
**processed and loop waiting for all prices while merging the information.
**Finally when all prices where processed
**returns the total count, the minimum, the maximum, the sum of all prices
 */
func PreciosYVentas(category string, ch chan ArgsAndResult, download Downloader) obtainedData {
	body := download(category)
	results := body["results"].([]interface{})
	total := GetTotalCount(&body)
	res := GetPricesAndSold(results)
	chanels := (total / 200) - 1
	// channs := make(map[int]chan obtainedData)
	responses1 := make(chan obtainedData)
	responses2 := make(chan obtainedData)
	responses3 := make(chan obtainedData)
	responses4 := make(chan obtainedData)
	responses5 := make(chan obtainedData)
	responses6 := make(chan obtainedData)
	responses7 := make(chan obtainedData)
	responses8 := make(chan obtainedData)
	responses9 := make(chan obtainedData)
	responses10 := make(chan obtainedData)
	responses11 := make(chan obtainedData)
	responses12 := make(chan obtainedData)
	responses13 := make(chan obtainedData)
	responses14 := make(chan obtainedData)
	responses15 := make(chan obtainedData)
	responses16 := make(chan obtainedData)
	responses17 := make(chan obtainedData)
	responses18 := make(chan obtainedData)
	responses19 := make(chan obtainedData)
	responses20 := make(chan obtainedData)
	responses21 := make(chan obtainedData)
	responses22 := make(chan obtainedData)
	responses23 := make(chan obtainedData)
	responses24 := make(chan obtainedData)
	responses25 := make(chan obtainedData)
	responses26 := make(chan obtainedData)
	responses27 := make(chan obtainedData)
	responses28 := make(chan obtainedData)
	responses29 := make(chan obtainedData)
	responses30 := make(chan obtainedData)
	responses31 := make(chan obtainedData)
	responses32 := make(chan obtainedData)
	responses33 := make(chan obtainedData)
	responses34 := make(chan obtainedData)
	responses35 := make(chan obtainedData)
	responses36 := make(chan obtainedData)
	responses37 := make(chan obtainedData)
	responses38 := make(chan obtainedData)
	responses39 := make(chan obtainedData)
	responses40 := make(chan obtainedData)
	//Start a Goroutine that would send in order all downloads waiting for any
	//Task worker free to download
	go func() {
		for c := 0; c < chanels; c += 20 {
			ch <- ArgsAndResult{responses1, category + "&offset=" + strconv.Itoa(200*(c+1)), download}
			ch <- ArgsAndResult{responses2, category + "&offset=" + strconv.Itoa(200*(c+2)), download}
			ch <- ArgsAndResult{responses3, category + "&offset=" + strconv.Itoa(200*(c+3)), download}
			ch <- ArgsAndResult{responses4, category + "&offset=" + strconv.Itoa(200*(c+4)), download}
			ch <- ArgsAndResult{responses5, category + "&offset=" + strconv.Itoa(200*(c+5)), download}
			ch <- ArgsAndResult{responses6, category + "&offset=" + strconv.Itoa(200*(c+6)), download}
			ch <- ArgsAndResult{responses7, category + "&offset=" + strconv.Itoa(200*(c+7)), download}
			ch <- ArgsAndResult{responses8, category + "&offset=" + strconv.Itoa(200*(c+8)), download}
			ch <- ArgsAndResult{responses9, category + "&offset=" + strconv.Itoa(200*(c+9)), download}
			ch <- ArgsAndResult{responses10, category + "&offset=" + strconv.Itoa(200*(c+10)), download}
			ch <- ArgsAndResult{responses11, category + "&offset=" + strconv.Itoa(200*(c+11)), download}
			ch <- ArgsAndResult{responses12, category + "&offset=" + strconv.Itoa(200*(c+12)), download}
			ch <- ArgsAndResult{responses13, category + "&offset=" + strconv.Itoa(200*(c+13)), download}
			ch <- ArgsAndResult{responses14, category + "&offset=" + strconv.Itoa(200*(c+14)), download}
			ch <- ArgsAndResult{responses15, category + "&offset=" + strconv.Itoa(200*(c+15)), download}
			ch <- ArgsAndResult{responses16, category + "&offset=" + strconv.Itoa(200*(c+16)), download}
			ch <- ArgsAndResult{responses17, category + "&offset=" + strconv.Itoa(200*(c+17)), download}
			ch <- ArgsAndResult{responses18, category + "&offset=" + strconv.Itoa(200*(c+18)), download}
			ch <- ArgsAndResult{responses19, category + "&offset=" + strconv.Itoa(200*(c+19)), download}
			ch <- ArgsAndResult{responses20, category + "&offset=" + strconv.Itoa(200*(c+20)), download}
			ch <- ArgsAndResult{responses21, category + "&offset=" + strconv.Itoa(200*(c+21)), download}
			ch <- ArgsAndResult{responses22, category + "&offset=" + strconv.Itoa(200*(c+22)), download}
			ch <- ArgsAndResult{responses23, category + "&offset=" + strconv.Itoa(200*(c+23)), download}
			ch <- ArgsAndResult{responses24, category + "&offset=" + strconv.Itoa(200*(c+24)), download}
			ch <- ArgsAndResult{responses25, category + "&offset=" + strconv.Itoa(200*(c+25)), download}
			ch <- ArgsAndResult{responses26, category + "&offset=" + strconv.Itoa(200*(c+26)), download}
			ch <- ArgsAndResult{responses27, category + "&offset=" + strconv.Itoa(200*(c+27)), download}
			ch <- ArgsAndResult{responses28, category + "&offset=" + strconv.Itoa(200*(c+28)), download}
			ch <- ArgsAndResult{responses29, category + "&offset=" + strconv.Itoa(200*(c+29)), download}
			ch <- ArgsAndResult{responses30, category + "&offset=" + strconv.Itoa(200*(c+30)), download}
			ch <- ArgsAndResult{responses31, category + "&offset=" + strconv.Itoa(200*(c+31)), download}
			ch <- ArgsAndResult{responses32, category + "&offset=" + strconv.Itoa(200*(c+32)), download}
			ch <- ArgsAndResult{responses33, category + "&offset=" + strconv.Itoa(200*(c+33)), download}
			ch <- ArgsAndResult{responses34, category + "&offset=" + strconv.Itoa(200*(c+34)), download}
			ch <- ArgsAndResult{responses35, category + "&offset=" + strconv.Itoa(200*(c+35)), download}
			ch <- ArgsAndResult{responses36, category + "&offset=" + strconv.Itoa(200*(c+36)), download}
			ch <- ArgsAndResult{responses37, category + "&offset=" + strconv.Itoa(200*(c+37)), download}
			ch <- ArgsAndResult{responses38, category + "&offset=" + strconv.Itoa(200*(c+38)), download}
			ch <- ArgsAndResult{responses39, category + "&offset=" + strconv.Itoa(200*(c+39)), download}
			ch <- ArgsAndResult{responses40, category + "&offset=" + strconv.Itoa(200*(c+40)), download}
		}
	}()

	/*
	**Wait for all channels to return and merge the prices information
	**then delete the channel key in the map to reduce the ammount of iterations
	 */
	//TODO: Evaluate if it would be better to set a sleeping time
	//before starting to loop again
	done := 0
	for done < chanels {
		select {
		case resi := <-responses1:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses2:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses3:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses4:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses5:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses6:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses7:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses8:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses9:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses10:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses11:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses12:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses13:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses14:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses15:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses16:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses17:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses18:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses19:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses20:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses21:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses22:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses23:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses24:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses25:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses26:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses27:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses28:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses29:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses30:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses31:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses32:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses33:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses34:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses35:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses36:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses37:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses38:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses39:
			MergeObainedData(&res, &resi)
			done++
		case resi := <-responses40:
			MergeObainedData(&res, &resi)
			done++
		default:
			continue
		}

	}
	return res
}
