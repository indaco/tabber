package tabber

import (
	"testing"
)

func TestEqual(t *testing.T) {
	tests := []struct {
		name string
		v1   int
		v2   int
		want bool
	}{
		{"Equal", 10, 10, true},
		{"NotEqual", 5, 10, false},
		{"NegativeValues", -5, -5, true},
		{"MixedSign", -5, 5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := equal(tt.v1, tt.v2)
			if got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntToString(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  string
	}{
		{"PositiveNumber", 10, "10"},
		{"Zero", 0, "0"},
		{"NegativeNumber", -5, "-5"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := intToString(tt.value)
			if got != tt.want {
				t.Errorf("IntToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrimmed(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"\"hello\"", "hello"},
		{"\"world\"", "world"},
		{"\"\"", ""},
		{"hello", "hello"}, // no quotes, no change expected
	}

	for _, test := range tests {
		result := trimmed(test.input)
		if result != test.expected {
			t.Errorf("Trimmed(%s) = %s; expected %s", test.input, result, test.expected)
		}
	}
}

func TestToSlug(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"New Resource", "new-resource"},
		{"Another_Resource", "another-resource"},
		{"Multiple  Spaces", "multiple--spaces"},
		{"Capital Letters", "capital-letters"},
		{"\"With Quotes\"", "with-quotes"}, // quotes should be removed
		{"", ""},                           // empty string should remain empty
	}

	for _, test := range tests {
		result := toSlug(test.input)
		if result != test.expected {
			t.Errorf("ToSlug(%s) = %s; expected %s", test.input, result, test.expected)
		}
	}
}
