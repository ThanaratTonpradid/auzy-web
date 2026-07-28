package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitizeString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Normal string",
			input:    "Hello World",
			expected: "Hello World",
		},
		{
			name:     "String with HTML",
			input:    "<script>alert('xss')</script>",
			expected: "&lt;script&gt;alert(&#39;xss&#39;)&lt;/script&gt;",
		},
		{
			name:     "String with whitespace",
			input:    "  Hello  ",
			expected: "Hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SanitizeString(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSanitizeUsername(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Valid username",
			input:    "User123",
			expected: "user123",
		},
		{
			name:     "Username with special chars",
			input:    "user@#$123",
			expected: "user123",
		},
		{
			name:     "Username with underscore and hyphen",
			input:    "user_name-123",
			expected: "user_name-123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SanitizeUsername(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSanitizeEmail(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Valid email",
			input:    "User@Example.COM",
			expected: "user@example.com",
		},
		{
			name:     "Invalid email",
			input:    "not-an-email",
			expected: "",
		},
		{
			name:     "Email with whitespace",
			input:    "  test@example.com  ",
			expected: "test@example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SanitizeEmail(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestIsValidPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		expected bool
	}{
		{
			name:     "Valid strong password",
			password: "StrongPass123!",
			expected: true,
		},
		{
			name:     "Too short",
			password: "Short1!",
			expected: false,
		},
		{
			name:     "No uppercase",
			password: "lowercase123!",
			expected: false,
		},
		{
			name:     "No lowercase",
			password: "UPPERCASE123!",
			expected: false,
		},
		{
			name:     "No digit",
			password: "NoDigits!@#",
			expected: false,
		},
		{
			name:     "No special char",
			password: "NoSpecial123",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidPassword(tt.password)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestRemoveHTMLTags(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Text with HTML tags",
			input:    "<p>Hello <strong>World</strong></p>",
			expected: "Hello World",
		},
		{
			name:     "Text without HTML tags",
			input:    "Plain text",
			expected: "Plain text",
		},
		{
			name:     "Script tag",
			input:    "<script>alert('xss')</script>Text",
			expected: "alert('xss')Text",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RemoveHTMLTags(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

