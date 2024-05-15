package tabber

// Variant represents the variant of the tab component.
type Variant string

// String returns the string representation of the Variant.
func (v Variant) String() string {
	return string(v)
}

// Placement represents where the tabs list should be positioned.
type Placement string

// String returns the string representation of the Placement.
func (p Placement) String() string {
	return string(p)
}

// Alignment represents where the tab list should be aligned.
type Alignment string

// String returns the string representation of the Alignement.
func (a Alignment) String() string {
	return string(a)
}

// Size represents the size of the tab component.
type Size string

// String returns the string representation of the Size.
func (s Size) String() string {
	return string(s)
}

type TabIcon struct {
	Value string
	Size  float32
}

func Icon(v string) *TabIcon {
	return &TabIcon{
		Value: v,
		Size:  1.25,
	}
}
