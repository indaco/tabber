package tabber

import (
	"testing"
)

func TestConfigBuilder(t *testing.T) {
	tests := []struct {
		name         string
		configParams func(*ConfigBuilder)
		expected     Config
	}{
		{
			name: "DefaultConfig",
			configParams: func(builder *ConfigBuilder) {
				// No additional configuration
			},
			expected: Config{
				ActiveTab: 1,
				Placement: Top,
				Variants:  []Variant{},
				Alignment: Start,
			},
		},
		{
			name: "CustomConfig",
			configParams: func(builder *ConfigBuilder) {
				builder.WithActiveTab(2).
					WithPlacement(Bottom).
					WithVariants(Accent, Boxed).
					WithAlignment(Center)
			},
			expected: Config{
				ActiveTab: 2,
				Placement: Bottom,
				Variants:  []Variant{Accent, Boxed},
				Alignment: Center,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewConfigBuilder()
			tt.configParams(builder)
			config := builder.Build()

			if config.ActiveTab != tt.expected.ActiveTab {
				t.Errorf("ActiveTab: got %d, want %d", config.ActiveTab, tt.expected.ActiveTab)
			}
			if config.Placement != tt.expected.Placement {
				t.Errorf("Placement: got %v, want %v", config.Placement, tt.expected.Placement)
			}
			if !equalVariants(config.Variants, tt.expected.Variants) {
				t.Errorf("Variants: got %v, want %v", config.Variants, tt.expected.Variants)
			}
			if config.Alignment != tt.expected.Alignment {
				t.Errorf("Alignment: got %v, want %v", config.Alignment, tt.expected.Alignment)
			}
		})
	}
}

func equalVariants(slice1, slice2 []Variant) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
