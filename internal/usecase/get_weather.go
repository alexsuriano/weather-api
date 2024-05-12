package usecase

import (
	"net/http"
	"strconv"

	externalapi "github.com/alexsuriano/weather-api/internal/external-api"
)

type GetWeatherHandler struct {
	Cep        string
	WeatherApi externalapi.WeatherApiHandler
	ViaCep     externalapi.ViaCepHandler
}

type Temperature struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

type ErrorResponse struct {
	HttpCode int    `json:"HttpCode"`
	Message  string `json:"message"`
}

func NewGetWeatherHandler(
	cep string,
	weatherApi externalapi.WeatherApiHandler,
	viaCep externalapi.ViaCepHandler,
) *GetWeatherHandler {
	return &GetWeatherHandler{
		Cep:        cep,
		WeatherApi: weatherApi,
		ViaCep:     viaCep,
	}
}

func (g *GetWeatherHandler) Execute() (*Temperature, *ErrorResponse) {
	if !g.ValidateCep() {
		return nil, &ErrorResponse{
			HttpCode: http.StatusUnprocessableEntity,
			Message:  "invalid zipcode",
		}
	}

	location, err := g.ViaCep.GetLocation(g.Cep)
	if err != nil {
		if err.Error() == "can not find zipcode" {
			return nil, &ErrorResponse{
				HttpCode: http.StatusNotFound,
				Message:  "can not find zipcode",
			}
		} else {
			return nil, &ErrorResponse{
				HttpCode: http.StatusInternalServerError,
				Message:  err.Error(),
			}
		}
	}

	weather, err := g.WeatherApi.GetTemperature(location.Localidade)
	if err != nil {
		return nil, &ErrorResponse{
			HttpCode: http.StatusBadRequest,
			Message:  err.Error(),
		}
	}

	return &Temperature{
		TempC: weather.Current.TempC,
		TempF: (weather.Current.TempC * 1.8) + 32,
		TempK: weather.Current.TempC + 273.15,
	}, nil
}

func (g *GetWeatherHandler) ValidateCep() bool {
	if g.Cep != "" {
		if len(g.Cep) == 8 {
			_, err := strconv.Atoi(g.Cep)
			return err == nil
		}
	}

	return false
}
