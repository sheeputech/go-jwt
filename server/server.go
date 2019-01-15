package main

import (
	"20181209sun-go-jwt/routers"
	"20181209sun-go-jwt/settings"
	"github.com/codegangsta/negroni"
	"github.com/rs/cors"
	"net/http"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://192.168.99.100"},
		AllowedHeaders: []string{"Authorization"},
	})

	settings.Init()
	router := routers.InitRoutes()

	n := negroni.Classic()
	n.Use(c)
	n.UseHandler(router)
	http.ListenAndServe(":9000", n)
}
