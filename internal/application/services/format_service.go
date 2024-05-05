package services

import "github.com/jeanSagaz/temperature-system-by-cep-gcp/internal/application/dto"

func FormatTemperatureService(localidade string, celsius float64) *dto.TemperatureResponse {
	return &dto.TemperatureResponse{
		Celsius:    celsius,
		Kelvin:     celsius + 273,
		Fahrenheit: (celsius * 1.8) + 32,
	}
}
