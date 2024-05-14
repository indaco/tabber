// Package tabber - Helpers for using tabber with `template/html`
package tabber

import (
	"context"
	"fmt"
	"html/template"
	"strings"

	"github.com/a-h/templ"
)

// HTMLGenerator provides functions for generating HTML code for tabber.
type HTMLGenerator struct{}

// NewHTMLGenerator creates a new instance of HTMLGenerator.
func NewHTMLGenerator() *HTMLGenerator {
	return &HTMLGenerator{}
}

// TabberCSSToGoHTML generates HTML code for the tabber CSS and returns it as a template.HTML.
func (g *HTMLGenerator) TabberCSSToGoHTML(ctx context.Context) (template.HTML, error) {
	tabberHTML, err := templ.ToGoHTML(context.Background(), TabberCSS(ctx))
	if err != nil {
		return "", fmt.Errorf("failed to generate tabber CSS: %v", err)
	}
	// Generate HTML for user-registered stylesheet variants
	var stylesBuilder strings.Builder
	configMap := GetConfigMapFromContext(ctx)
	for _, config := range configMap.Data {
		for _, variant := range config.Variants {
			if _, ok := registeredTabberVariants[variant.String()]; ok {
				style, err := templ.ToGoHTML(ctx, UseVariant(variant.String()))
				if err != nil {
					return "", fmt.Errorf("failed to generate tabber CSS for variant %s: %v", variant.String(), err)
				}
				stylesBuilder.WriteString(string(style))
			}
		}
	}

	// Combine TabberCSS HTML and variant styles HTML
	joinedHTML := tabberHTML + template.HTML(stylesBuilder.String())
	return joinedHTML, nil
}

// TabberJSToGoHTML generates HTML code for the tabber JS and returns it as a template.HTML.
func (g *HTMLGenerator) TabberJSToGoHTML(configMap *ConfigMap) (template.HTML, error) {
	html, err := templ.ToGoHTML(context.Background(), TabberJS(configMap))
	if err != nil {
		return "", fmt.Errorf("failed to generate tabber JS: %v", err)
	}
	return html, nil
}

// Render generates HTML code for displaying the tabber and returns it as a template.HTML.
func (g *HTMLGenerator) Render(ctx context.Context, comp templ.Component) (template.HTML, error) {
	// Generate HTML code for displaying the toast.
	html, err := templ.ToGoHTML(ctx, comp)
	if err != nil {
		return "", fmt.Errorf("failed to generate tabber HTML: %v", err)
	}
	return html, nil
}
