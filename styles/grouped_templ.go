// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.680
package styles

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

// Variant Name: Grouped
func TabberGroupedCss() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style>\n    [data-variant=\"grouped\"] .gtb_tab_list {\n      position: relative;\n      justify-content: center;\n      isolation: isolate;\n\n      button {\n        transition: background-color 0.3s ease;\n      }\n\n      button:focus,\n      button:focus-visible {\n        --gtb-tab-ring-width: 1px;\n        z-index: 1;\n      }\n\n      button:is(.active, [aria-selected=\"true\"]) {\n        --gtb-tab-bg: var(--gtb-tab-bg-active);\n      }\n\n      button:is(.active, [aria-selected=\"true\"])active:focus,\n      button:is(.active, [aria-selected=\"true\"]):focus-visible {\n        --gtb-tab-bg-focus: var(--gtb-base-light);\n        --gtb-tab-ring-width: 1px;\n        z-index: 1;\n      }\n\n      @media screen and (width > 320px) {\n        & {\n          grid-template-columns: repeat(auto-fit, minmax(100px, max-content));\n          --gtb-tab-list-gap: 0;\n\n          button:first-of-type {\n            --gtb-tab-border-tl-radius: 100vh;\n            --gtb-tab-border-tr-radius: 0;\n            --gtb-tab-border-bl-radius: 100vh;\n            --gtb-tab-border-br-radius: 0;\n          }\n\n          button:last-of-type {\n            --gtb-tab-border-tl-radius: 0;\n            --gtb-tab-border-tr-radius: 100vh;\n            --gtb-tab-border-bl-radius: 0;\n            --gtb-tab-border-br-radius: 100vh;\n          }\n        }\n      }\n\n      /* Media Queries */\n      @media screen and (width > 640px) {\n        [data-placement=\"left\"] &,\n        [data-placement=\"right\"] & {\n          grid-template-columns: auto;\n          align-content: space-around;\n          transform-origin: center center;\n          margin-bottom: 0;\n          padding: 0;\n          padding-block: 3em;\n          writing-mode: horizontal-tb;\n        }\n\n        [data-placement=\"left\"] & {\n          button {\n            justify-content: center;\n            transform: rotate(-90deg);\n          }\n\n          button:first-of-type {\n            --gtb-tab-border-tl-radius: 0;\n            --gtb-tab-border-tr-radius: 100vh;\n            --gtb-tab-border-bl-radius: 0;\n            --gtb-tab-border-br-radius: 100vh;\n          }\n\n          button:last-of-type {\n            --gtb-tab-border-tl-radius: 100vh;\n            --gtb-tab-border-tr-radius: 0;\n            --gtb-tab-border-bl-radius: 100vh;\n            --gtb-tab-border-br-radius: 0;\n          }\n        }\n\n        [data-placement=\"right\"] & {\n          button {\n            justify-content: center;\n            transform: rotate(90deg);\n          }\n\n          button:first-of-type {\n            --gtb-tab-border-tl-radius: 100vh;\n            --gtb-tab-border-tr-radius: 0;\n            --gtb-tab-border-bl-radius: 100vh;\n            --gtb-tab-border-br-radius: 0;\n          }\n\n          button:last-of-type {\n            --gtb-tab-border-tl-radius: 0;\n            --gtb-tab-border-tr-radius: 100vh;\n            --gtb-tab-border-bl-radius: 0;\n            --gtb-tab-border-br-radius: 100vh;\n          }\n        }\n      }\n\n      @media screen and (width > 720px) {\n        [data-placement=\"left\"] &,\n        [data-placement=\"right\"] & {\n          padding-block: 1em;\n        }\n      }\n\n      @media screen and (width > 1280px) {\n        [data-placement=\"left\"] &,\n        [data-placement=\"right\"] & {\n          padding-block: 2.5em;\n        }\n      }\n    }\n\n    /* Media Feature - prefers-color-scheme */\n    @media (prefers-color-scheme: light) {\n      :root {\n        [data-variant=\"grouped\"] {\n          --gtb-tab-bg: var(--gtb-base-lightest);\n          --gtb-tab-bg-hover: var(--gtb-base-lighter);\n          --gtb-tab-bg-focus: var(--gtb-base-light);\n          --gtb-tab-bg-active: var(--gtb-base-light);\n        }\n      }\n    }\n\n    @media (prefers-color-scheme: dark) {\n      :root {\n        [data-variant=\"grouped\"] {\n          --gtb-tab-bg: var(--gtb-base-darker);\n          --gtb-tab-bg-hover: var(--gtb-base);\n          --gtb-tab-bg-focus: var(--gtb-base-dark);\n          --gtb-tab-bg-active: var(--gtb-base-dark);\n        }\n      }\n    }\n  </style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}