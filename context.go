package tabber

import (
	"context"
)

// Define the context key type.
type contextKey string

func (k contextKey) String() string {
	return string(k)
}

// GetConfigMapFromContext retrieves the tab configuration from the context.
func GetConfigMapFromContext(ctx context.Context) *ConfigMap {
	// go-staticcheck(SA1029):not use built-in type string as key for value, define your own type.
	if opts, ok := ctx.Value(ConfigContextKey).(*ConfigMap); ok {
		return opts
	}
	// When echo.Context, to get data the key must be a string.
	if opts, ok := ctx.Value(ConfigContextKey.String()).(*ConfigMap); ok {
		return opts
	}
	return &ConfigMap{
		Data: make(map[string]Config),
	}
}

// GetConfigFromContextById retrieves the tab configuration from the context.
func GetConfigFromContextById(ctx context.Context, id string) *Config {
	if config, ok := GetConfigMapFromContext(ctx).Get(id); ok {
		return &config
	}
	return &Config{}
}

// GetVariantsFromContextById retrieves the configured tab variant from the context.
func GetVariantsFromContextById(ctx context.Context, id string) []Variant {
	if config, ok := GetConfigMapFromContext(ctx).Get(id); ok {
		return config.Variants
	}
	return []Variant{}
}

// getFirstVariantAsStringFromContextById retrieves the configured tab variant from the context.
func getFirstVariantAsStringFromContextById(ctx context.Context, id string) string {
	if len(GetVariantsFromContextById(ctx, id)) > 0 {
		return GetVariantsFromContextById(ctx, id)[0].String()
	}
	return Headless.String()
}

// getPlacementFromContextById retrieves the configured tabs list placement from the context.
func getPlacementFromContextById(ctx context.Context, id string) Placement {
	if config, ok := GetConfigMapFromContext(ctx).Get(id); ok {
		return config.Placement
	}
	return DefaultConfig().Placement
}

// getPlacementAsStringFromContextById retrieves the configured tabs list placement from the context.
func getPlacementAsStringFromContextById(ctx context.Context, id string) string {
	return getPlacementFromContextById(ctx, id).String()
}

// getAlignmentFromContextById retrieves the configured tab list alignment from the context.
func getAlignmentFromContextById(ctx context.Context, id string) Alignment {
	if config, ok := GetConfigMapFromContext(ctx).Get(id); ok {
		return config.Alignment
	}
	return DefaultConfig().Alignment
}

// getAlignmentAsStringFromContextById retrieves the configured tab list alignment from the context as string.
func getAlignmentAsStringFromContextById(ctx context.Context, id string) string {
	return getAlignmentFromContextById(ctx, id).String()
}
