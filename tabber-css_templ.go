// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package tabber

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"context"
	"github.com/indaco/tabber/styles"
)

func TabberCSS(ctx context.Context) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<style>\n    *,\n    *::before,\n    *::after {\n      box-sizing: border-box;\n      border: 0;\n      border-style: solid;\n    }\n\n    :root {\n      /* - Color Shades - */\n      --gtb-color-space: oklab;\n      --gtb-base: hsl(215deg 16.3% 46.9%);\n\n      --gtb-base-10: color-mix(\n        in var(--gtb-color-space),\n        var(--gtb-base) 10%,\n        white\n      );\n\n      --gtb-base-20: color-mix(\n        in var(--gtb-color-space),\n        var(--gtb-base) 30%,\n        white\n      );\n\n      --gtb-base-30: color-mix(\n        in var(--gtb-color-space),\n        var(--gtb-base) 50%,\n        white\n      );\n\n      --gtb-base-40: color-mix(\n        in var(--gtb-color-space),\n        var(--gtb-base) 70%,\n        white\n      );\n\n      --gtb-base-50: var(--gtb-base);\n\n      --gtb-base-60: color-mix(\n        in var(--gtb-color-space),\n        var(--gtb-base) 70%,\n        black\n      );\n\n      --gtb-base-70: color-mix(\n        in var(--gtb-color-space),\n        var(--gtb-base) 50%,\n        black\n      );\n\n      --gtb-base-80: color-mix(\n        in var(--gtb-color-space),\n        var(--gtb-base) 30%,\n        black\n      );\n\n      --gtb-base-90: color-mix(\n        in var(--gtb-color-space),\n        var(--gtb-base) 10%,\n        black\n      );\n\n      --gtb-base-lightest: var(--gtb-base-10);\n      --gtb-base-lighter: var(--gtb-base-20);\n      --gtb-base-light: var(--gtb-base-40);\n      --gtb-base-dark: var(--gtb-base-60);\n      --gtb-base-darker: var(--gtb-base-80);\n      --gtb-base-darkest: var(--gtb-base-90);\n\n      /* ******************************************** *\n       * Tabber Component\n       * ******************************************** */\n\n      /* - TabContainer - */\n      --gtb-root-width: 100%;\n      --gtb-root-max-width: fit-content;\n      --gtb-root-my: 0;\n      --gtb-root-mx: 0;\n      --gtb-root-py: 0;\n      --gtb-root-px: 0;\n      --gtb-root-border-width: 0;\n      --gtb-root-border-style: solid;\n      --gtb-root-border-color: transparent;\n      --gtb-root-border-radius: 0;\n\n      /* - TabList - */\n      --gtb-tab-list-width: auto;\n      --gtb-tab-list-max-width: calc(100% + 20%);\n      --gtb-tab-list-alignment: flex-start;\n      --gtb-tab-list-my: 0;\n      --gtb-tab-list-mx: 0;\n      --gtb-tab-list-py: 0;\n      --gtb-tab-list-px: 0;\n      --gtb-tab-list-gap: 0.5em;\n      --gtb-tab-list-bg: var(--gtb-bg);\n      --gtb-tab-list-border-t-width: 0;\n      --gtb-tab-list-border-t-style: solid;\n      --gtb-tab-list-border-t-color: var(--gtb-border-color);\n      --gtb-tab-list-border-r-width: 0;\n      --gtb-tab-list-border-r-style: solid;\n      --gtb-tab-list-border-r-color: var(--gtb-border-color);\n      --gtb-tab-list-border-b-width: 0;\n      --gtb-tab-list-border-b-style: solid;\n      --gtb-tab-list-border-b-color: var(--gtb-border-color);\n      --gtb-tab-list-border-l-width: 0;\n      --gtb-tab-list-border-l-style: solid;\n      --gtb-tab-list-border-l-color: var(--gtb-border-color);\n      --gtb-tab-list-border-tl-radius: 0;\n      --gtb-tab-list-border-tr-radius: 0;\n      --gtb-tab-list-border-bl-radius: 0;\n      --gtb-tab-list-border-br-radius: 0;\n\n      /* - Tab - */\n      --gtb-tab-py: 0.75em;\n      --gtb-tab-px: 1em;\n      --gtb-tab-color: var(--gtb-color);\n      --gtb-tab-color-hover: var(--gtb-color-hover);\n      --gtb-tab-color-focus: var(--gtb-color-focus);\n      --gtb-tab-color-active: var(--gtb-color-active);\n      --gtb-tab-font-family: inherit;\n      --gtb-tab-bg: transparent;\n      --gtb-tab-bg-hover: var(--gtb-bg-hover);\n      --gtb-tab-bg-focus: var(--gtb-bg-focus);\n      --gtb-tab-bg-active: var(--gtb-bg-active);\n      --gtb-tab-border-t-width: 0;\n      --gtb-tab-border-t-style: solid;\n      --gtb-tab-border-t-color: var(--gtb-border-color);\n      --gtb-tab-border-r-width: 0;\n      --gtb-tab-border-r-style: solid;\n      --gtb-tab-border-r-color: var(--gtb-border-color);\n      --gtb-tab-border-b-width: 0;\n      --gtb-tab-border-b-style: solid;\n      --gtb-tab-border-b-color: var(--gtb-border-color);\n      --gtb-tab-border-l-width: 0;\n      --gtb-tab-border-l-style: solid;\n      --gtb-tab-border-l-color: var(--gtb-border-color);\n      --gtb-tab-border-t-color-hover: var(--gtb-tab-border-t-color);\n      --gtb-tab-border-r-color-hover: var(--gtb-tab-border-r-color);\n      --gtb-tab-border-b-color-hover: var(--gtb-tab-border-b-color);\n      --gtb-tab-border-l-color-hover: var(--gtb-tab-border-l-color);\n      --gtb-tab-border-tl-radius: 0;\n      --gtb-tab-border-tr-radius: 0;\n      --gtb-tab-border-bl-radius: 0;\n      --gtb-tab-border-br-radius: 0;\n      --gtb-tab-ring-width: 0;\n      --gtb-tab-ring-style: solid;\n      --gtb-tab-ring-offset: 0.125rem;\n      --gtb-tab-icon-gap-x: 0.3em;\n      --gtb-tab-icon-gap-y: 0.3em;\n      --gtb-tab-icon-alignment: row;\n\n      /* - TabPanel - */\n      --gtb-panel-width: 100%;\n      --gtb-panel-max-width: 680px;\n      --gtb-panel-my: 0;\n      --gtb-panel-mx: 0;\n      --gtb-panel-py: 1.25em;\n      --gtb-panel-px: 1.25em;\n      --gtb-panel-color: var(--gtb-color);\n      --gtb-panel-bg: var(--gtb-bg);\n      --gtb-panel-border-t-width: 0;\n      --gtb-panel-border-t-style: solid;\n      --gtb-panel-border-t-color: var(--gtb-border-color);\n      --gtb-panel-border-r-width: 0;\n      --gtb-panel-border-r-style: solid;\n      --gtb-panel-border-r-color: var(--gtb-border-color);\n      --gtb-panel-border-b-width: 0;\n      --gtb-panel-border-b-style: solid;\n      --gtb-panel-border-b-color: var(--gtb-border-color);\n      --gtb-panel-border-l-width: 0;\n      --gtb-panel-border-l-style: solid;\n      --gtb-panel-border-l-color: var(--gtb-border-color);\n      --gtb-panel-border-tl-radius: 0;\n      --gtb-panel-border-tr-radius: 0;\n      --gtb-panel-border-bl-radius: 0;\n      --gtb-panel-border-br-radius: 0;\n    }\n\n    @media (prefers-color-scheme: light) {\n      :root {\n        /* Simplified set of css custom props to make themes */\n        --gtb-color: var(--gtb-base-darker);\n        --gtb-color-hover: var(--gtb-base-darker);\n        --gtb-color-focus: var(--gtb-base-darker);\n        --gtb-color-active: var(--gtb-base-darkest);\n\n        --gtb-bg: inherit;\n        --gtb-bg-hover: inherit;\n        --gtb-bg-focus: inherit;\n        --gtb-bg-active: inherit;\n\n        --gtb-border-color: var(--gtb-base-lighter);\n        --gtb-border-active: var(--gtb-base);\n\n        --gtb-nav-bg: var(--gtb-base-lighter);\n\n        --gtb-tab-ring-color: var(--gtb-base-light);\n      }\n    }\n\n    @media (prefers-color-scheme: dark) {\n      :root {\n        /* Simplified set of css custom props to make themes */\n        --gtb-color: var(--gtb-base-lighter);\n        --gtb-color-hover: var(--gtb-base-lighter);\n        --gtb-color-focus: var(--gtb-base-lighter);\n        --gtb-color-active: var(--gtb-base-lightest);\n\n        --gtb-bg: inherit;\n        --gtb-bg-hover: inherit;\n        --gtb-bg-focus: inherit;\n        --gtb-bg-active: inherit;\n\n        --gtb-border-color: var(--gtb-base-light);\n        --gtb-border-active: var(--gtb-base);\n\n        --gtb-nav-bg: var(--gtb-base-darker);\n\n        --gtb-tab-ring-color: var(--gtb-base);\n      }\n    }\n  </style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = styles.TabberHeadlessCss().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, config := range GetConfigMapFromContext(ctx).Data {
			for _, variant := range config.Variants {
				if _, ok := registeredTabberVariants[variant.String()]; ok {
					templ_7745c5c3_Err = UseVariant(variant.String()).Render(ctx, templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
			}
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
