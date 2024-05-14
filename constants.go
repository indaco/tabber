package tabber

// ConfigContextKey is a context key for the tab component configurations.
var ConfigContextKey = contextKey("tabber-config-ctx")

// Placement defines where the tab list will appear on the screen.
var (
	Top    Placement = "top"
	Bottom Placement = "bottom"
	Left   Placement = "left"
	Right  Placement = "right"
)

// Alignement defines where the tab list will be aligned.
var (
	Start  Alignment = "start"
	Center Alignment = "center"
	End    Alignment = "end"
)

// Variant defines different visual variants for the tab component.
var (
	Headless   Variant = "headless"   // Headless (unstyled) variant.
	Accent     Variant = "accent"     // Ghost variant.
	Boxed      Variant = "boxed"      // Boxed variant.
	Grouped    Variant = "grouped"    // Grouped variant.
	Pills      Variant = "pills"      // Pills variant.
	Rounded    Variant = "rounded"    // Rounded variant.
	Underlined Variant = "underlined" // Underlined variant.
)
