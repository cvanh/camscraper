package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main(){
	getCatagory()
	fmt.Println("hallo")
}


// get list of products by catagory number(125 for nut and cams on klimwinkel)
func getCatagory(){
	client := &http.Client{}

	reqbody := []byte(`{
	"alternative_options_photos": true,
	"category_id": 125,
	"brands_filter": [],
	"stock_filter": [],
	"specs_filter": [],
	"sort": "default"
	`)

	req, err := http.NewRequest(http.MethodPost, "https://www.klimwinkel.nl/api/queryProducts", bytes.NewBuffer(reqbody))
	if err != nil {
		fmt.Println("error while building req")
	}

	// we want a json response other wise we will get html 
	req.Header.Set("Content-Type","application/json;charset=utf-8")

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("error while fetching catagories")
	}

	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := io.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))

}
