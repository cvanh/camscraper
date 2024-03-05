package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/IBM/sarama"
)

// TODO auto generated please cleanup
type KlimwinkelCatResponse struct {
	Products struct {
		CurrentPage int `json:"current_page"`
		Data        []struct {
			ID                       int           `json:"id"`
			Brand                    string        `json:"brand"`
			BrandDeliveryDaysMin     int           `json:"brand_delivery_days_min"`
			BrandDeliveryDaysMax     int           `json:"brand_delivery_days_max"`
			Name                     string        `json:"name"`
			FullName                 string        `json:"full_name"`
			Usp1                     string        `json:"usp1"`
			Usp2                     string        `json:"usp2"`
			Usp3                     string        `json:"usp3"`
			Slug                     string        `json:"slug"`
			URI                      string        `json:"uri"`
			ShortDescription         string        `json:"short_description"`
			Price                    float64       `json:"price"`
			PriceGross               float64       `json:"price_gross"`
			OldPrice                 float32       `json:"old_price"`
			OldPriceGross            float32       `json:"old_price_gross"`
			StockStatus              int           `json:"stock_status"`
			Photo                    string        `json:"photo"`
			PhotoFull                string        `json:"photo_full"`
			BestSelling              bool          `json:"best_selling"`
			PoID                     int           `json:"po_id"`
			OnRequest                bool          `json:"on_request"`
			OptionsCount             int           `json:"options_count"`
			AlternativesCount        int           `json:"alternatives_count"`
			AlternativeOptionsPhotos []interface{} `json:"alternative_options_photos"`
		} `json:"data"`
		FirstPageURL string `json:"first_page_url"`
		From         int    `json:"from"`
		LastPage     int    `json:"last_page"`
		LastPageURL  string `json:"last_page_url"`
		Links        []struct {
			URL    interface{} `json:"url"`
			Label  string      `json:"label"`
			Active bool        `json:"active"`
		} `json:"links"`
		NextPageURL interface{} `json:"next_page_url"`
		Path        string      `json:"path"`
		PerPage     int         `json:"per_page"`
		PrevPageURL interface{} `json:"prev_page_url"`
		To          int         `json:"to"`
		Total       int         `json:"total"`
	} `json:"products"`
	MinPrice      float64 `json:"min_price"`
	MaxPrice      float64 `json:"max_price"`
	MinPriceGross float64 `json:"min_price_gross"`
	MaxPriceGross float64 `json:"max_price_gross"`
	StockCount    struct {
		Num1 int `json:"1"`
		Num2 int `json:"2"`
	} `json:"stock_count"`
	Specs []struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
		Vals  []struct {
			Val struct {
				ID    int    `json:"id"`
				Title string `json:"title"`
			} `json:"val"`
			Unit  interface{} `json:"unit"`
			Count int         `json:"count"`
		} `json:"vals"`
	} `json:"specs"`
	Brands []struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Count int    `json:"count"`
	} `json:"brands"`
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
		Product struct {
			ID                   int    `json:"id"`
			Brand                string `json:"brand"`
			BrandDeliveryDaysMin int    `json:"brand_delivery_days_min"`
			BrandDeliveryDaysMax int    `json:"brand_delivery_days_max"`
			Name                 string `json:"name"`
			MetaTitle            string `json:"meta_title"`
			MetaDescription      string `json:"meta_description"`
			FullName             string `json:"full_name"`
			Description          string `json:"description"`
			Photos               []struct {
				ID        int    `json:"id"`
				Type      string `json:"type"`
				Value     string `json:"value"`
				Thumbnail string `json:"thumbnail"`
				Width     string `json:"width"`
				Height    string `json:"height"`
				Filename  string `json:"filename,omitempty"`
			} `json:"photos"`
			Replacement        interface{}   `json:"replacement"`
			Favorite           interface{}   `json:"favorite"`
			BestSelling        bool          `json:"best_selling"`
			OnSale             bool          `json:"on_sale"`
			CoverPhoto         interface{}   `json:"cover_photo"`
			OnRequest          bool          `json:"on_request"`
			Slug               string        `json:"slug"`
			Usp1               string        `json:"usp1"`
			Usp2               string        `json:"usp2"`
			Usp3               string        `json:"usp3"`
			ContentBlocks      []interface{} `json:"content_blocks"`
			ExtraFields        []interface{} `json:"extra_fields"`
			AlternateLangHrefs []struct {
				LanguageID int    `json:"language_id"`
				URI        string `json:"uri"`
			} `json:"alternate_lang_hrefs"`
			ExtraGuarantee interface{} `json:"extra_guarantee"`
		} `json:"product"`
		Breadcrumbs []struct {
			Title string `json:"title"`
			URI   string `json:"uri"`
		} `json:"breadcrumbs"`
		Pdfs             []interface{} `json:"pdfs"`
		ReferenceNumbers []string      `json:"reference_numbers"`
		Specifications   []struct {
			ID    int    `json:"id"`
			Title string `json:"title"`
			Value string `json:"value"`
		} `json:"specifications"`
		ProsAndCons []struct {
			Type int    `json:"type"`
			Text string `json:"text"`
		} `json:"pros_and_cons"`
		Questions    []interface{} `json:"questions"`
		Alternatives []struct {
			ID                   int         `json:"id"`
			Brand                string      `json:"brand"`
			BrandDeliveryDaysMin int         `json:"brand_delivery_days_min"`
			BrandDeliveryDaysMax int         `json:"brand_delivery_days_max"`
			Name                 string      `json:"name"`
			FullName             string      `json:"full_name"`
			Usp1                 string      `json:"usp1"`
			Usp2                 string      `json:"usp2"`
			Usp3                 string      `json:"usp3"`
			Slug                 string      `json:"slug"`
			URI                  string      `json:"uri"`
			ShortDescription     string      `json:"short_description"`
			Price                float64     `json:"price"`
			PriceGross           float64     `json:"price_gross"`
			OldPrice             interface{} `json:"old_price"`
			OldPriceGross        interface{} `json:"old_price_gross"`
			StockStatus          int         `json:"stock_status"`
			Photo                string      `json:"photo"`
			PhotoFull            string      `json:"photo_full"`
			BestSelling          bool        `json:"best_selling"`
			PoID                 int         `json:"po_id"`
			OnRequest            bool        `json:"on_request"`
			OptionsCount         int         `json:"options_count"`
			AlternativesCount    interface{} `json:"alternatives_count"`
		} `json:"alternatives"`
		Pages []struct {
			ID                 int         `json:"id"`
			Parent             int         `json:"parent"`
			RootParent         interface{} `json:"root_parent"`
			Type               int         `json:"type"`
			TypeCode           string      `json:"type_code"`
			URI                string      `json:"uri"`
			URIPage            string      `json:"uri_page"`
			AlternateLangHrefs interface{} `json:"alternate_lang_hrefs"`
			Title              string      `json:"title"`
			MetaTitle          string      `json:"meta_title"`
			MetaDesc           string      `json:"meta_desc"`
			Excerpt            string      `json:"excerpt"`
			Media              []struct {
				ID        int    `json:"id"`
				Type      string `json:"type"`
				Value     string `json:"value"`
				Thumbnail string `json:"thumbnail"`
				Width     string `json:"width"`
				Height    string `json:"height"`
				Filename  string `json:"filename"`
			} `json:"media"`
			Breadcrumbs  interface{} `json:"breadcrumbs"`
			CustomFields interface{} `json:"custom_fields"`
			Schema       struct {
				Context   string `json:"@context"`
				Type      string `json:"@type"`
				Headline  string `json:"headline"`
				Publisher struct {
					Type string `json:"@type"`
					Name string `json:"name"`
				} `json:"publisher"`
				Author struct {
					Type string `json:"@type"`
					Name string `json:"name"`
				} `json:"author"`
				Image        string    `json:"image"`
				DateModified time.Time `json:"dateModified"`
			} `json:"schema"`
		} `json:"pages"`
		DefaultAttrs []interface{} `json:"default_attrs"`
		Reviews      struct {
			Items         []interface{} `json:"items"`
			Count         int           `json:"count"`
			AverageRating interface{}   `json:"average_rating"`
		} `json:"reviews"`
	} `json:"data"`
}

var brokers = []string{"127.0.0.1:9092"}

func main() {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)

	if err != nil {
		log.Fatalf("error while connecting to kafka %v",err)
	}


	products := getCatagory()
	for _, v := range products.Products.Data {
		_ = v
		data := getProduct(v.ID)
		log.Println("%v", data.Data.Schema)

		productSchema, _ := json.Marshal(data.Data.Schema)

		msg := &sarama.ProducerMessage{
			Topic:     "scraped_data",
			Partition: -1,
			Value:	sarama.StringEncoder(productSchema), 
		}

		_,_,err := producer.SendMessage(msg)
		

		log.Println("%v",err)
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
