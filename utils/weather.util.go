package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

type owmWeather struct {
	Description string `json:"description"`
}

type owmMain struct {
	Temp      float32 `json:"temp"`
	FeelsLike float32 `json:"feels_like"`
	Humidity  float32 `json:"humidity"`
}

type weatherRes struct {
	Weather []owmWeather `json:"weather"`
	Main    owmMain      `json:"main"`
}

func GetWeather(loc string) weatherRes {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create a base url
	baseUrl, err := url.Parse("https://api.openweathermap.org/data/2.5/weather")
	if err != nil {
		log.Fatalf("Failed to parse url: %v", err)
	}

	// Create query params
	params := url.Values{}
	params.Add("q", loc)
	params.Add("units", "imperial")
	params.Add("appid", os.Getenv("WEATHER_API_KEY"))

	// Add queries to the url
	baseUrl.RawQuery = params.Encode()

	// Make request to Open Weather Map (owm)
	res, err := http.Get(baseUrl.String())
	if err != nil {
		log.Fatalf("Failed to make weather request: %v", err)
	}

	defer res.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Failed to read the response body: %v", err)
	}

	var w weatherRes // Store the unmarshalled response
	err = json.Unmarshal(body, &w)
	if err != nil {
		log.Fatalf("Failed to unmarshal the body: %v", err)
	}

	return w
}
