package analyzer

import (
	"testing"

	"github.com/Vladroon22/linters-check/config"
	"github.com/Vladroon22/linters-check/pkg/analyzer"
)

func TestCheckLowerCase(t *testing.T) {
	tests := []struct {
		name     string
		msg      string
		expected bool
	}{
		{
			name:     "Upper case",
			msg:      "Hello",
			expected: false,
		},
		{
			name:     "lower case",
			msg:      "hello",
			expected: true,
		},
	}

	cfg := config.NewConfig()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := analyzer.CheckLowerCase(cfg, tt.msg)

			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestCheckEnglishOnly(t *testing.T) {
	tests := []struct {
		name     string
		msg      string
		expected bool
	}{
		{
			name:     "russian",
			msg:      "Привет мир",
			expected: false,
		},
		{
			name:     "english",
			msg:      "hello world",
			expected: true,
		},
		{
			name:     "egnlish + russian",
			msg:      "hello привет",
			expected: false,
		},
		{
			name:     "eegnlish + mandarin",
			msg:      "hello 世界",
			expected: false,
		},
	}
	cfg := config.NewConfig()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := analyzer.CheckEnglishOnly(cfg, tt.msg)

			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestCheckSpecialChars(t *testing.T) {
	tests := []struct {
		name     string
		msg      string
		expected bool
	}{
		{
			name:     "check disabled",
			msg:      "hello @#$%",
			expected: false,
		},
		{
			name:     "empty message",
			msg:      "",
			expected: true,
		},
		{
			name:     "valid message",
			msg:      "hello world 123",
			expected: true,
		},
		{
			name:     "invalid special char @",
			msg:      "hello @ world",
			expected: true,
		},
		{
			name:     "invalid special char #",
			msg:      "test #hash",
			expected: true,
		},
		{
			name:     "invalid special char $",
			msg:      "price $100",
			expected: false,
		},
		{
			name:     "emoji smiley",
			msg:      "hello 😊 world",
			expected: false,
		},
		{
			name:     "emoji heart",
			msg:      "i ❤️ go",
			expected: false,
		},
		{
			name:     "emoji thumbs up",
			msg:      "good job 👍",
			expected: false,
		},
		{
			name:     "multiple emojis",
			msg:      "test 🎉🚀🔥",
			expected: false,
		},
	}
	cfg := config.NewConfig()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := analyzer.CheckSpecialChars(cfg, tt.msg)

			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestCheckSensitiveData(t *testing.T) {
	tests := []struct {
		name     string
		msg      string
		expected bool
	}{
		{
			name:     "check disabled",
			msg:      "user password is 123",
			expected: false,
		},
		{
			name:     "empty message",
			msg:      "",
			expected: true,
		},
		{
			name:     "contains password",
			msg:      "user password is 12345",
			expected: false,
		},
		{
			name:     "contains passwd",
			msg:      "passwd=12345",
			expected: false,
		},
		{
			name:     "contains secret",
			msg:      "api secret key",
			expected: false,
		},
		{
			name:     "contains token",
			msg:      "auth token expired",
			expected: false,
		},
		{
			name:     "contains credit card",
			msg:      "credit card number: 4111-112311-12311111-1111",
			expected: false,
		},
		{
			name:     "case insensitive",
			msg:      "USER PASSWORD IS SECRET",
			expected: false,
		},
		{
			name:     "multiple sensitive keywords",
			msg:      "password and token are sensitive",
			expected: false,
		},
	}
	cfg := config.NewConfig()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := analyzer.CheckSensitiveData(cfg, tt.msg)

			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}

}

func TestCheckEmoji(t *testing.T) {
	tests := []struct {
		name     string
		msg      string
		expected bool
	}{
		{
			name:     "transport range 1F680-1F6FF",
			msg:      "test 🚀🚁🚂🚃",
			expected: false,
		},
		{
			name:     "cyrillic characters",
			msg:      "привет мир",
			expected: true,
		},
		{
			name:     "emoji thumb up",
			msg:      "test 👍🏻",
			expected: false,
		},
		{
			name:     "family emoji",
			msg:      "test 👨‍👩‍👧",
			expected: false,
		},
	}

	cfg := config.NewConfig()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := analyzer.CheckEmoji(cfg, tt.msg)

			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
