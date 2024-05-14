package tabber

import (
	"github.com/a-h/templ"
	"github.com/indaco/tabber/styles"
)

// TabberVariant is an interface for different tabber variants.
type TabberVariant interface {
	GetClassName() templ.Component
}

// AccentVariant represents the accent variant of the tabber.
type AccentVariant struct{}

// GetClassName returns the CSS class name for the accent variant of the tabber.
func (hv AccentVariant) GetClassName() templ.Component {
	return styles.TabberAccentCss()
}

// BoxedVariant represents the boxed variant of the tabber.
type BoxedVariant struct{}

// GetClassName returns the CSS class name for the boxed variant of the tabber.
func (bv BoxedVariant) GetClassName() templ.Component {
	return styles.TabberBoxedCss()
}

// GroupedVariant represents the grouped variant of the tabber.
type GroupedVariant struct{}

// GetClassName returns the CSS class name for the grouped variant of the tabber.
func (gv GroupedVariant) GetClassName() templ.Component {
	return styles.TabberGroupedCss()
}

// PillsVariant represents the pills variant of the tabber.
type PillsVariant struct{}

// GetClassName returns the CSS class name for the pills variant of the tabber.
func (pv PillsVariant) GetClassName() templ.Component {
	return styles.TabberPillsCss()
}

// RoundedVariant represents the rounded variant of the tabber.
type RoundedVariant struct{}

// GetClassName returns the CSS class name for the rounded variant of the tabber.
func (rv RoundedVariant) GetClassName() templ.Component {
	return styles.TabberRoundedCss()
}

// UnderlinedVariant represents the underlined variant of the tabber.
type UnderlinedVariant struct{}

// GetClassName returns the CSS class name for the underlined variant of the tabber.
func (uv UnderlinedVariant) GetClassName() templ.Component {
	return styles.TabberUnderlinedCss()
}

// UseVariant returns the appropriate tabber variant component based on the provided variant name.
func UseVariant(variant string) templ.Component {
	if ev, ok := registeredTabberVariants[variant]; ok {
		return ev.GetClassName()
	}
	return nil
}
