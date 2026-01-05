/*

	Expected Response format example

	{
		"location": {
			"name": "London",
			"region": "City of London, Greater London",
			"country": "United Kingdom",
			"lat": 51.5171,
			"lon": -0.1062,
			"tz_id": "Europe/London",
			"localtime_epoch": 1767633260,
			"localtime": "2026-01-05 17:14"
		},
		"current": {
			"last_updated_epoch": 1767632400,
			"last_updated": "2026-01-05 17:00",
			"temp_c": 0.1,
			"temp_f": 32.2,
			"is_day": 0,
			"condition": {
				"text": "Clear",
				"icon": "//cdn.weatherapi.com/weather/64x64/night/113.png",
				"code": 1000
			},
			"wind_mph": 8.1,
			"wind_kph": 13.0,
			"wind_degree": 318,
			"wind_dir": "NW",
			"pressure_mb": 1015.0,
			"pressure_in": 29.97,
			"precip_mm": 0.0,
			"precip_in": 0.0,
			"humidity": 69,
			"cloud": 0,
			"feelslike_c": -3.9,
			"feelslike_f": 25.0,
			"windchill_c": -4.0,
			"windchill_f": 24.8,
			"heatindex_c": 0.0,
			"heatindex_f": 32.0,
			"dewpoint_c": -6.1,
			"dewpoint_f": 21.1,
			"vis_km": 10.0,
			"vis_miles": 6.0,
			"uv": 0.0,
			"gust_mph": 13.1,
			"gust_kph": 21.1,
			"air_quality": {
				"co": 258.85,
				"no2": 55.55,
				"o3": 10.0,
				"so2": 4.85,
				"pm2_5": 9.95,
				"pm10": 14.25,
				"us-epa-index": 1,
				"gb-defra-index": 1
			},
			"short_rad": 0,
			"diff_rad": 0,
			"dni": 0,
			"gti": 0
		}
	}
*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const apiKey = "94b62b37083e4a23bab171303260501"
const aqi = true

type response struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzID           string  `json:"tz_id"`
		LocalTime      string  `json:"localtime"`
		LocalTimeEpoch int64   `json:"localtime_epoch"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int64   `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph         float64 `json:"wind_mph"`
		WindKph         float64 `json:"wind_kph"`
		WindDegree      int     `json:"wind_degree"`
		WindDir         string  `json:"wind_dir"`
		PressureMb      float64 `json:"pressure_mb"`
		PressureIn      float64 `json:"pressure_in"`
		PrecipitationMm float64 `json:"precip_mm"`
		PrecipitationIn float64 `json:"precip_in"`
		Humidity        int     `json:"humidity"`
		Cloud           int     `json:"cloud"`
		FeelsLikeC      float64 `json:"feelslike_c"`
		FeelsLikeF      float64 `json:"feelslike_f"`
		WindChillC      float64 `json:"windchill_c"`
		WindChillF      float64 `json:"windchill_f"`
		HeatIndexC      float64 `json:"heatindex_c"`
		HeatIndexF      float64 `json:"heatindex_f"`
		DewPointC       float64 `json:"dewpoint_c"`
		DewPointF       float64 `json:"dewpoint_f"`
		VisKm           float64 `json:"vis_km"`
		VisMiles        float64 `json:"vis_miles"`
		Uv              float64 `json:"uv"`
		GustMph         float64 `json:"gust_mph"`
		GustKph         float64 `json:"gust_kph"`
		AirQuality      struct {
			Co           float64 `json:"co"`
			NO2          float64 `json:"no2"`
			O3           float64 `json:"o3"`
			SO2          float64 `json:"so2"`
			PM2_5        float64 `json:"pm2_5"`
			PM10         float64 `json:"pm10"`
			USEPAIndex   float64 `json:"us-epa-index"`
			GBDefraIndex float64 `json:"gb-defra-index"`
		} `json:"air_quality"`
		ShortRad float64 `json:"short_rad"`
		DiffRad  float64 `json:"diff_rad"`
		Dni      float64 `json:"dni"`
		Gti      float64 `json:"gti"`
	} `json:"current"`
}

func main() {

	startTime := time.Now()

	cities := []string{"London", "Tokyo", "Paris", "Toronto"}

	ch := make(chan string)
	var wg sync.WaitGroup

	for _, city := range cities {
		wg.Add(1)
		go fetchWeather(city, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Print(result)
	}

	fmt.Printf("Time taken: %s\n", time.Since(startTime))
}

func fetchWeather(city string, ch chan<- string, wg *sync.WaitGroup) response {
	var data response

	defer wg.Done()

	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=%s", apiKey, city, strconv.FormatBool(aqi))
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error fetching data")
		return data
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println(err)
		return data
	}

	ch <- fmt.Sprintf("City: %s, Temperature in Celsius: %v\n", city, data.Current.TempC)

	return data
}

// Time taken avg : 400ms to 200ms
