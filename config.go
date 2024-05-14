package tabber

type Config struct {
	ActiveTab int
	Placement Placement
	Variants  []Variant
	Alignment Alignment
}

// ConfigBuilder is a builder for creating Config objects.
type ConfigBuilder struct {
	config Config
}

type ConfigMap struct {
	Data map[string]Config
}

// DefaultConfig returns the default configuration.
func DefaultConfig() Config {
	return Config{
		ActiveTab: 1,
		Placement: Top,
		Variants:  []Variant{},
		Alignment: Start,
	}
}

// NewConfigBuilder creates a new ConfigBuilder with default values.
func NewConfigBuilder() *ConfigBuilder {
	return &ConfigBuilder{
		config: DefaultConfig(),
	}
}

// WithActiveTab sets the active tab index.
func (b *ConfigBuilder) WithActiveTab(activeTab int) *ConfigBuilder {
	b.config.ActiveTab = activeTab
	return b
}

// WithPlacement sets the placement of the tabs list.
func (b *ConfigBuilder) WithPlacement(placement Placement) *ConfigBuilder {
	b.config.Placement = placement
	return b
}

// WithVariant configures a single style variant.
func (b *ConfigBuilder) WithVariant(variant Variant) *ConfigBuilder {
	if v, exists := registeredTabberVariants[variant.String()]; !exists {
		// If not registered, register it
		RegisterVariant(variant, v)
	}
	b.config.Variants = append(b.config.Variants, variant)
	return b
}

// WithVariants sets the variants for the tab.
func (b *ConfigBuilder) WithVariants(variants ...Variant) *ConfigBuilder {
	for _, variant := range variants {
		b.WithVariant(variant)
	}
	return b
}

// WithAlignment sets the alignment of the tab list.
func (b *ConfigBuilder) WithAlignment(alignment Alignment) *ConfigBuilder {
	b.config.Alignment = alignment
	return b
}

// Build constructs the Config object.
func (b *ConfigBuilder) Build() Config {
	return b.config
}

// Initialize a new ConfigMap instance
func NewConfigMap() *ConfigMap {
	return &ConfigMap{
		Data: make(map[string]Config),
	}
}

// Add adds a configuration to the map.
func (m *ConfigMap) Add(key string, config Config) {
	m.Data[key] = config
}

// Get retrieves a configuration from the map.
func (m *ConfigMap) Get(key string) (Config, bool) {
	config, ok := m.Data[key]
	return config, ok
}
