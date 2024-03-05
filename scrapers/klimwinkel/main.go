package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/IBM/sarama"
)

type KlimwinkelCatResponse struct {
	Products struct {
		Data []struct {
			ID int `json:"id"`
		} `json:"data"`
	}
}
type KlimwinkelProductResponse struct {
	Data struct {
		Schema struct {
			Context string `json:"@context"`
			Type    string `json:"@type"`
			Brand   struct {
				Type string `json:"@type"`
				Name string `json:"name"`
			} `json:"brand"`
			Sku         string `json:"sku"`
			Mpn         string `json:"mpn"`
			Image       string `json:"image"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Offers      struct {
				Type          string  `json:"@type"`
				OfferCount    int     `json:"offerCount"`
				LowPrice      float64 `json:"lowPrice"`
				HighPrice     float64 `json:"highPrice"`
				PriceCurrency string  `json:"priceCurrency"`
				Offers        []struct {
					Type          string  `json:"@type"`
					URL           string  `json:"url"`
					Availability  string  `json:"availability"`
					Price         float64 `json:"price"`
					PriceCurrency string  `json:"priceCurrency"`
					Gtin          string  `json:"gtin"`
				} `json:"offers"`
				Availability string `json:"availability"`
			} `json:"offers"`
			AggregateRating struct {
				Type        string `json:"@type"`
				BestRating  int    `json:"bestRating"`
				RatingValue string `json:"ratingValue"`
				ReviewCount int    `json:"reviewCount"`
			} `json:"aggregateRating"`
			Review struct {
				Context string `json:"@context"`
				Type    string `json:"@type"`
				Author  struct {
					Type string `json:"@type"`
					Name string `json:"name"`
				} `json:"author"`
				DatePublished string `json:"datePublished"`
				ReviewBody    string `json:"reviewBody"`
				ReviewRating  struct {
					Type        string `json:"@type"`
					BestRating  int    `json:"bestRating"`
					RatingValue int    `json:"ratingValue"`
				} `json:"reviewRating"`
			} `json:"review"`
		} `json:"schema"`
	} `json:"data"`
}

// TODO should be changable when ran
var brokers = []string{"127.0.0.1:9092"}

func main() {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)

	if err != nil {
		log.Fatalf("error while connecting to kafka %v", err)
	}

	products := getCatagory()
	for _, v := range products.Products.Data {
		_ = v
		data := getProduct(v.ID)

		productSchema, _ := json.Marshal(data.Data.Schema)

		msg := &sarama.ProducerMessage{
			Topic:     "scraped_data",
			Headers:   []sarama.RecordHeader{
				sarama.RecordHeader{
					Key:   []byte("source"),
					Value: []byte("klimwinkel.nl"),
				},
			},
			Value:     sarama.StringEncoder(productSchema),
		}

		_, _, err := producer.SendMessage(msg)
		if err != nil {
			log.Printf("error while sending message, got: %v", err)
		} 
	}

}

func getProduct(id int) KlimwinkelProductResponse {
	client := &http.Client{}

	url := fmt.Sprint("https://www.klimwinkel.nl/api/products/", id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	// TODO make shure these wont expire
	req.Header.Set("Cookie", "fzshoprest_session=eyJpdiI6Ik1sR0tnQlBOU0wwL2ZBdHU5bm1NZXc9PSIsInZhbHVlIjoiZm1XZE5MeGE5Q1hTRjdnQTFzSU1IQ0hDTDlEUUtZUzg5WTh6VnhUSE9xS3RKdDVUK1I0YjltbU1rOE91TmcvOUJpVnNydDd3SVY2TnlHUmgxdnhUdzd5OG1saDRiVUJwTWcreTQ3ZHFtQU9hZHM3a08xM1BPdjRBNm8zUkpURk8iLCJtYWMiOiIxZjkyYzgzNTM5YTk4MzRlZWQ1MjYwYjYwODU2OTAyZDQ3NmIyZmRkYmVhNTE1MjQ3ZTY0YjdhZjlmNjYzMTZhIiwidGFnIjoiIn0%253D; i18n_redirected=nl; auth.strategy=local")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var Product KlimwinkelProductResponse
	json.Unmarshal(bodyText, &Product)

	return Product

}

// get list of products by catagory number(125 for nut and cams on klimwinkel)
func getCatagory() KlimwinkelCatResponse {
	client := &http.Client{}
	data := strings.NewReader(`{
	"alternative_options_photos": true,
	"category_id": 125,
	"brands_filter": [],
	"stock_filter": [],
	"specs_filter": [],
	"sort": "default"
	}`)

	req, err := http.NewRequest("POST", "https://www.klimwinkel.nl/api/queryProducts", data)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	// TODO make shure these wont expire
	req.Header.Set("Cookie", "fzshoprest_session=eyJpdiI6Ik1sR0tnQlBOU0wwL2ZBdHU5bm1NZXc9PSIsInZhbHVlIjoiZm1XZE5MeGE5Q1hTRjdnQTFzSU1IQ0hDTDlEUUtZUzg5WTh6VnhUSE9xS3RKdDVUK1I0YjltbU1rOE91TmcvOUJpVnNydDd3SVY2TnlHUmgxdnhUdzd5OG1saDRiVUJwTWcreTQ3ZHFtQU9hZHM3a08xM1BPdjRBNm8zUkpURk8iLCJtYWMiOiIxZjkyYzgzNTM5YTk4MzRlZWQ1MjYwYjYwODU2OTAyZDQ3NmIyZmRkYmVhNTE1MjQ3ZTY0YjdhZjlmNjYzMTZhIiwidGFnIjoiIn0%253D; i18n_redirected=nl; auth.strategy=local")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var Products KlimwinkelCatResponse
	json.Unmarshal(bodyText, &Products)

	return Products
}
