package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/indaco/tabber"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	firstTabberConfig := tabber.NewConfigBuilder().WithVariant(tabber.Boxed).Build()
	secondTabberConfig := tabber.NewConfigBuilder().WithVariant(tabber.Underlined).Build()

	configMap := tabber.NewConfigMap()
	configMap.Add("tabs-1", firstTabberConfig)
	configMap.Add("tabs-2", secondTabberConfig)

	ctx := context.WithValue(r.Context(), tabber.ConfigContextKey, configMap)
	err := HomePage().Render(ctx, w)
	if err != nil {
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", HandleHome)

	port := ":3300"
	log.Printf("Listening on %s", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Printf("failed to start server: %v", err)
		os.Exit(1)
	}
}
