package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

func query(city string) (*weatherData, error) {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data := new(weatherData)
	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, err
	}
	return data, nil
}

func weather(w http.ResponseWriter, r *http.Request) {
	city := strings.SplitN(r.URL.Path, "/", 3)[2]
	data, err := query(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
}

func main() {
	http.HandleFunc("/weather/", weather)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
