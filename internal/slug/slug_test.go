package slug

import (
	"testing"
)

func TestMakeSlug(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name:     "single word",
			input:    []string{"Hello"},
			expected: "hello",
		},
		{
			name:     "multiple words",
			input:    []string{"Hello", "World"},
			expected: "hello-world",
		},
		{
			name:     "empty input",
			input:    []string{},
			expected: "",
		},
		{
			name:     "words with special characters",
			input:    []string{"Hello!", "World@2023"},
			expected: "hello!-world@2023",
		},
		{
			name:     "words with mixed case",
			input:    []string{"HeLLo", "WorLD"},
			expected: "hello-world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MakeSlug(tt.input)
			if result != tt.expected {
				t.Errorf("MakeSlug(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCleanWords(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "single word",
			input:    []string{"Hello"},
			expected: []string{"Hello"},
		},
		{
			name:     "multiple words",
			input:    []string{"Hello", "World"},
			expected: []string{"Hello", "World"},
		},
		{
			name:     "words with special characters",
			input:    []string{"Hello!", "World@2023"},
			expected: []string{"Hello-", "World-2023"},
		},
		{
			name:     "words with spaces",
			input:    []string{"Hello ", " World"},
			expected: []string{"Hello", "World"},
		},
		{
			name:     "words with numbers",
			input:    []string{"Hello123", "World456"},
			expected: []string{"Hello123", "World456"},
		},
		{
			name:     "words with symbols",
			input:    []string{"Hello#", "World$"},
			expected: []string{"Hello-", "World-"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CleanWords(tt.input)
			for i, v := range result {
				if v != tt.expected[i] {
					t.Errorf("CleanWords(%v) = %v; want %v", tt.input, result, tt.expected)
					break
				}
			}
		})
	}
}
