package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/indaco/tabber"
)

type PageData struct {
	Tabber TabberComponent
}

type TabberComponent struct {
	CSS  template.HTML
	HTML template.HTML
	JS   template.HTML
}

// HandleHome is the handler function for the home page "/"
func HandleHome(w http.ResponseWriter, r *http.Request) {

	tabConfig := tabber.NewConfigBuilder().WithVariant(tabber.Underlined).Build()

	configMap := tabber.NewConfigMap()
	configMap.Add("demo", tabConfig)
	ctx := context.WithValue(r.Context(), tabber.ConfigContextKey, configMap)

	tabHTMLGenerator := tabber.NewHTMLGenerator()

	// Render the needed CSS for tab component as template.HTML
	tabCSS, err := tabHTMLGenerator.TabberCSSToGoHTML(ctx)
	if err != nil {
		log.Printf("failed to render template: %v", err)
	}

	// Render the tab component into a template.HTML
	tabHtml, err := tabHTMLGenerator.Render(ctx, TabDemo())
	if err != nil {
		log.Printf("failed to render template: %v", err)
	}

	// Render the needed JS for tab component as template.HTML
	tabJS, err := tabHTMLGenerator.TabberJSToGoHTML(configMap)
	if err != nil {
		log.Printf("failed to render template: %v", err)
	}

	data := PageData{
		Tabber: TabberComponent{
			CSS:  tabCSS,
			HTML: tabHtml,
			JS:   tabJS,
		},
	}

	// Parse the HTML template
	tmpl := template.Must(template.New("index").Parse(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>tabber - template/html</title>

			<!-- inject component styles -->
			{{ .Tabber.CSS }}

			<!-- styles -->
			<style type="text/css">
				.centered {
					width: fit-content;
					height: fit-content;
					margin: auto;
				}

				p {
					line-height: 1.25;
				}
			</style>
		</head>
		<body>
			<div class="centered">
				<!-- display the tab -->
				{{ .Tabber.HTML }}
			</div>
			<!-- inject component javascript -->
			{{ .Tabber.JS }}
		</body>
		</html>
	`))

	// Execute the template with the provided data and write the output to the response writer
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "failed to render template", http.StatusInternalServerError)
		log.Printf("failed to render template: %v", err)
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
