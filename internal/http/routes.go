package http

import "github.com/alexsuriano/weather-api/internal/http/webserver"

func SetRoutes(server *webserver.Webserver) {
	server.AddHandler("/temp", GetTemp)
	server.AddHandler("/healthCheck", HealthCheck)
}
