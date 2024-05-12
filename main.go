package main

import (
	"github.com/alexsuriano/weather-api/internal/http"
	"github.com/alexsuriano/weather-api/internal/http/webserver"
)

func main() {
	server := webserver.NewWebserver(":8080")
	http.SetRoutes(server)
	server.Start()
}
