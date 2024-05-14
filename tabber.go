package tabber

var (
	registeredTabberVariants map[string]TabberVariant
)

// RegisterVariant registers a tabber variant by name.
func RegisterVariant(name Variant, variant TabberVariant) {
	registeredTabberVariants[name.String()] = variant
}

func init() {
	registeredTabberVariants = make(map[string]TabberVariant)

	RegisterVariant(Accent, AccentVariant{})
	RegisterVariant(Boxed, BoxedVariant{})
	RegisterVariant(Grouped, GroupedVariant{})
	RegisterVariant(Pills, PillsVariant{})
	RegisterVariant(Rounded, RoundedVariant{})
	RegisterVariant(Underlined, UnderlinedVariant{})
}
