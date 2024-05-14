package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"unicode"

	"github.com/indaco/tabber"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	enabledVariants := []tabber.Variant{
		tabber.Accent,
		tabber.Boxed,
		tabber.Grouped,
		tabber.Pills,
		tabber.Rounded,
		tabber.Underlined,
	}

	config := tabber.NewConfigBuilder().WithVariants(enabledVariants...).Build()

	configMap := tabber.NewConfigMap()
	configMap.Add("custom-color", config)
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

func toTitle(s string) string {
	if len(s) < 1 {
		return s
	}
	rs := []rune(s)
	rs[0] = unicode.ToTitle(rs[0])
	return string(rs)
}
