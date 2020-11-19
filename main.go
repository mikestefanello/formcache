package main

import (
	"fmt"
	"net/http"

	"github.com/mikestefanello/formcache/cache"
	"github.com/mikestefanello/formcache/handlers"
	"github.com/mikestefanello/formcache/router"
	"github.com/rs/zerolog/log"
)

func main() {
	// Create a new cache
	c := cache.NewCache()

	// Create an HTTP handler
	handler := handlers.NewHTTPHandler(c)

	// Load the router
	r := router.NewRouter(handler)

	// Start the server
	addr := fmt.Sprintf("localhost:9000")
	log.Info().Str("on", addr).Msg("Server started")
	err := http.ListenAndServe(addr, r)
	if err != nil {
		log.Fatal().Err(err).Msg("Server terminated")
	}
}
