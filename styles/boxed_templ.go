// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package styles

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

// Variant Name: Boxed
func TabberBoxedCss() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<style>\n    :root {\n      [data-variant=\"boxed\"] {\n        --gtb-tab-list-gap: 0.5rem;\n        --gtb-tab-bg-active: var(--_body-bg);\n        --gtb-tab-bg-focus: var(--gtb-panel-bg);\n\n        --gtb-panel-border-t-width: 1px;\n        --gtb-panel-border-r-width: 1px;\n        --gtb-panel-border-b-width: 1px;\n        --gtb-panel-border-l-width: 1px;\n      }\n    }\n\n    /* - Tab List - */\n    [data-variant=\"boxed\"] .gtb_tab_list {\n      margin-bottom: 0.75em;\n\n      [data-placement=\"top\"] & {\n        button:is(.active, [aria-selected=\"true\"]) {\n          --gtb-tab-border-t-width: 3px;\n          --gtb-tab-border-t-color: var(--gtb-border-active);\n        }\n\n        button:is(.active, [aria-selected=\"true\"]):focus,\n        button:is(.active, [aria-selected=\"true\"]):focus-visible {\n          --gtb-tab-border-t-color: var(--gtb-base);\n          margin-bottom: -0.05em;\n        }\n      }\n\n      [data-placement=\"bottom\"] & {\n        margin-top: 0.75em;\n        margin-bottom: 0;\n\n        button:is(.active, [aria-selected=\"true\"]) {\n          --gtb-tab-border-b-width: 3px;\n          --gtb-tab-border-b-color: var(--gtb-border-active);\n        }\n\n        button:is(.active, [aria-selected=\"true\"]):focus,\n        button:is(.active, [aria-selected=\"true\"]):focus-visible {\n          --gtb-tab-border-b-color: var(--gtb-base);\n          margin-top: -0.05em;\n        }\n      }\n\n      [data-placement=\"left\"] & {\n        button:is(.active, [aria-selected=\"true\"]) {\n          --gtb-tab-border-l-width: 3px;\n          --gtb-tab-border-l-color: var(--gtb-border-active);\n        }\n\n        button:is(.active, [aria-selected=\"true\"]):focus,\n        button:is(.active, [aria-selected=\"true\"]):focus-visible {\n          --gtb-tab-border-l-color: var(--gtb-base);\n        }\n      }\n\n      [data-placement=\"right\"] & {\n        button:is(.active, [aria-selected=\"true\"]) {\n          --gtb-tab-border-r-width: 3px;\n          --gtb-tab-border-r-color: var(--gtb-border-active);\n        }\n\n        button:is(.active, [aria-selected=\"true\"]):focus,\n        button:is(.active, [aria-selected=\"true\"]):focus-visible {\n          --gtb-tab-border-r-color: var(--gtb-base);\n        }\n      }\n\n      /* Media Queries */\n      @media screen and (width > 480px) {\n        [data-placement=\"top\"] & {\n\n          margin-bottom: 0;\n\n          button:is(.active, [aria-selected=\"true\"]) {\n            --gtb-tab-border-r-width: 1px;\n            --gtb-tab-border-r-color: var(--gtb-border-color);\n\n            --gtb-tab-border-b-width: 1px;\n            --gtb-tab-border-b-color: var(--_body-bg, hsl(0deg 0% 100%));\n\n            --gtb-tab-border-l-width: 1px;\n            --gtb-tab-border-l-color: var(--gtb-border-color);\n\n            margin-bottom: -0.08em;\n          }\n        }\n\n        [data-placement=\"bottom\"] & {\n\n          margin-top: 0;\n\n          button:is(.active, [aria-selected=\"true\"]) {\n            --gtb-tab-border-t-width: 1px;\n            --gtb-tab-border-t-color: var(--_body-bg, hsl(0deg 0% 100%));\n\n            --gtb-tab-border-r-width: 1px;\n            --gtb-tab-border-r-color: var(--gtb-border-color);\n\n            --gtb-tab-border-l-width: 1px;\n            --gtb-tab-border-l-color: var(--gtb-border-color);\n\n            margin-top: -0.2em;\n          }\n        }\n\n        [data-placement=\"left\"] & {\n          button:is(.active, [aria-selected=\"true\"]) {\n            --gtb-tab-border-t-width: 1px;\n            --gtb-tab-border-t-color: var(--gtb-border-color);\n\n            --gtb-tab-border-r-width: 1px;\n            --gtb-tab-border-r-color: var(--gtb-border-color);\n\n            margin-bottom: -0.2em;\n          }\n        }\n\n        [data-placement=\"right\"] & {\n          button:is(.active, [aria-selected=\"true\"]) {\n            --gtb-tab-border-t-width: 1px;\n            --gtb-tab-border-t-color: var(--gtb-border-color);\n\n            --gtb-tab-border-l-width: 1px;\n            --gtb-tab-border-l-color: var(--gtb-border-color);\n\n            margin-bottom: -0.2em;\n          }\n        }\n      }\n\n      @media screen and (width > 620px) {\n        [data-placement=\"top\"] & {\n          button:is(.active, [aria-selected=\"true\"]) {\n            margin-bottom: -0.1em;\n          }\n        }\n\n        [data-placement=\"bottom\"] & {\n          button:is(.active, [aria-selected=\"true\"]) {\n            margin-top: -0.06em;\n          }\n        }\n      }\n\n      @media screen and (width > 640px) {\n        [data-placement=\"top\"] & {\n          margin-bottom: 0;\n        }\n\n        [data-placement=\"bottom\"] & {\n          margin-top: 0;\n        }\n\n        [data-placement=\"left\"] & {\n          --gtb-tab-list-border-t-width: 1px;\n          --gtb-tab-list-border-b-width: 1px;\n          --gtb-tab-list-border-l-width: 1px;\n\n          margin-bottom: 0;\n          padding-block: 2rem;\n          writing-mode: horizontal-tb;\n\n          button:is(.active, [aria-selected=\"true\"]) {\n            --gtb-tab-border-t-width: 0;\n            --gtb-tab-border-r-width: 0;\n          }\n        }\n\n        [data-placement=\"right\"] & {\n          --gtb-tab-list-border-t-width: 1px;\n          --gtb-tab-list-border-r-width: 1px;\n          --gtb-tab-list-border-b-width: 1px;\n\n          margin-bottom: 0;\n          padding-block: 2rem;\n          writing-mode: horizontal-tb;\n\n          button:is(.active, [aria-selected=\"true\"]) {\n            --gtb-tab-border-t-width: 0;\n            --gtb-tab-border-l-width: 0;\n          }\n        }\n      }\n    }\n\n    /* Media Feature - prefers-color-scheme */\n    @media (prefers-color-scheme: light) {\n      [data-variant=\"boxed\"] {\n        --gtb-tab-bg-hover: var(--gtb-base-lightest);\n      }\n    }\n\n    @media (prefers-color-scheme: dark) {\n      [data-variant=\"boxed\"] {\n        --gtb-tab-bg-hover: var(--gtb-base-dark);\n      }\n    }\n  </style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
