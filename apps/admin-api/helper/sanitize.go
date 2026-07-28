package helper

import (
	"html"
	"regexp"
	"strings"
)

// SanitizeString removes potentially dangerous characters from a string
func SanitizeString(input string) string {
	// Trim whitespace
	sanitized := strings.TrimSpace(input)
	
	// Escape HTML special characters
	sanitized = html.EscapeString(sanitized)
	
	return sanitized
}

// SanitizeUsername sanitizes username input (alphanumeric, underscore, hyphen only)
func SanitizeUsername(username string) string {
	// Remove any character that's not alphanumeric, underscore, or hyphen
	re := regexp.MustCompile(`[^a-zA-Z0-9_-]`)
	sanitized := re.ReplaceAllString(username, "")
	
	// Trim and lowercase
	sanitized = strings.TrimSpace(sanitized)
	sanitized = strings.ToLower(sanitized)
	
	return sanitized
}

// SanitizeEmail sanitizes email input
func SanitizeEmail(email string) string {
	// Trim and lowercase
	sanitized := strings.TrimSpace(email)
	sanitized = strings.ToLower(sanitized)
	
	// Basic email format validation
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(sanitized) {
		return ""
	}
	
	return sanitized
}

// RemoveHTMLTags removes all HTML tags from a string
func RemoveHTMLTags(input string) string {
	// Remove HTML tags
	re := regexp.MustCompile(`<[^>]*>`)
	return re.ReplaceAllString(input, "")
}

// SanitizeSQL prevents basic SQL injection by escaping dangerous characters
func SanitizeSQL(input string) string {
	// Replace single quotes with two single quotes (SQL escaping)
	sanitized := strings.ReplaceAll(input, "'", "''")
	
	// Remove null bytes
	sanitized = strings.ReplaceAll(sanitized, "\x00", "")
	
	return sanitized
}

// ValidateAndSanitizeInput validates and sanitizes common input types
func ValidateAndSanitizeInput(input string, inputType string) (string, bool) {
	switch inputType {
	case "username":
		sanitized := SanitizeUsername(input)
		isValid := len(sanitized) >= 3 && len(sanitized) <= 50
		return sanitized, isValid
		
	case "email":
		sanitized := SanitizeEmail(input)
		isValid := sanitized != ""
		return sanitized, isValid
		
	case "text":
		sanitized := SanitizeString(input)
		isValid := len(sanitized) <= 1000
		return sanitized, isValid
		
	case "name":
		sanitized := SanitizeString(input)
		isValid := len(sanitized) >= 1 && len(sanitized) <= 100
		return sanitized, isValid
		
	default:
		return SanitizeString(input), true
	}
}

// IsValidPassword validates password strength
func IsValidPassword(password string) bool {
	// At least 8 characters
	if len(password) < 8 {
		return false
	}
	
	// Contains at least one uppercase letter
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	// Contains at least one lowercase letter
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	// Contains at least one digit
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
	// Contains at least one special character
	hasSpecial := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`).MatchString(password)
	
	return hasUpper && hasLower && hasDigit && hasSpecial
}

