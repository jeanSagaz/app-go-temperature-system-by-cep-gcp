package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/jeanSagaz/temperature-system-by-cep-gcp/internal/application/dto"
)

func GetWeatherApiService(city string) (*dto.WeatherResponse, error) {
	key := "8c5d11c25d764ebdba311342240502"
	c := strings.ReplaceAll(city, " ", "&nbsp;")
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", key, c)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var weatherResponse dto.WeatherResponse
	err = json.Unmarshal(body, &weatherResponse)
	if err != nil {
		return nil, err
	}

	return &weatherResponse, nil
}
