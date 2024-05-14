package tabber

import (
	"context"
	"testing"
)

// setup initializes the test environment before each test.
func setup(t *testing.T) context.Context {
	configMap := NewConfigMap()
	configMap.Add("demo", Config{ActiveTab: 1, Variants: []Variant{Rounded}})
	return context.WithValue(context.Background(), ConfigContextKey, configMap)
}

func TestGetConfigFromContextById(t *testing.T) {
	// Arrange
	ctx := setup(t)
	want := &Config{ActiveTab: 1}

	// Act
	got := GetConfigFromContextById(ctx, "demo")

	// Assert
	if got.ActiveTab != want.ActiveTab {
		t.Errorf("GetConfigFromContextById() = %v, want %v", got, want)
	}
}

func TestGetFirstVariantAsStringFromContextById(t *testing.T) {
	// Arrange
	ctx := setup(t)
	want := "rounded"

	// Act
	got := getFirstVariantAsStringFromContextById(ctx, "demo")

	// Assert
	if got != want {
		t.Errorf("getFirstVariantAsStringFromContextById() = %v, want %v", got, want)
	}
}
