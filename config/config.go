package config

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	path = "config.yaml"
)

type Config struct {
	CheckLowercase     bool     `yaml:"check_lower_case" json:"check_lower_case"`
	CheckEnglish       bool     `yaml:"check_english" json:"check_english"`
	CheckSpecialChars  bool     `yaml:"check_special_sym" json:"check_special_sym"`
	CheckSensitiveData bool     `yaml:"check_sensitive" json:"check_sensitive"`
	SensitiveRegex     []string `yaml:"sens_regex" json:"sens_regex"`
	SensitiveKeywords  []string `yaml:"sens_keyword" json:"sens_keyword"`
}

func NewConfig() *Config {
	cnf, err := loadConfig()
	if err == nil {
		return cnf
	}

	return &Config{
		CheckLowercase:     true,
		CheckEnglish:       true,
		CheckSpecialChars:  true,
		CheckSensitiveData: true,
		SensitiveRegex:     patterns,
		SensitiveKeywords:  keywords,
	}
}

var patterns = []string{
	"(?i)password",
	"(?i)passwd",
	"(?i)pwd",
	"(?i)token",
	"(?i)jwt",
	"(?i)api_?key",
	"(?i)apikey",
	"(?i)secret",
	"(?i)private.?key",
	"(?i)client.?secret",
	"(?i)credential",
	"(?i)auth",
	"(?i)bearer",
	"credit.?card",
}

var keywords = []string{
	"password", "passwd", "pwd",
	"secret", "token", "api_key", "apikey",
	"key", "auth", "authorization",
	"credit", "card", "cvv",
	"private", "cert", "certificate",
	"credit",
	"social security",
}

func loadConfig() (*Config, error) {
	ext := strings.Split(path, ".")
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	var config Config
	switch ext[1] {
	case "json":
		if err := json.Unmarshal(data, &config); err != nil {
			return nil, fmt.Errorf("error parsing JSON: %w", err)
		}
	case "yaml":
		config = loadYaml()
	}

	return &config, nil
}

func loadYaml() Config {
	settings := make(map[string]any)

	cfg := Config{}

	if val, ok := settings["check_lower_case"].(bool); ok {
		cfg.CheckLowercase = val
	}
	if val, ok := settings["check_english"].(bool); ok {
		cfg.CheckEnglish = val
	}
	if val, ok := settings["check_special_sym"].(bool); ok {
		cfg.CheckSpecialChars = val
	}
	if val, ok := settings["check_sensitive"].(bool); ok {
		cfg.CheckSensitiveData = val
	}

	if keywords, ok := settings["sens_keyword"].([]interface{}); ok {
		for _, kw := range keywords {
			if str, ok := kw.(string); ok {
				cfg.SensitiveKeywords = append(cfg.SensitiveKeywords, str)
			}
		}
	}

	if regexes, ok := settings["sens_regex"].([]interface{}); ok {
		for _, re := range regexes {
			if str, ok := re.(string); ok {
				if compiled, err := regexp.Compile(str); err == nil {
					cfg.SensitiveRegex = append(cfg.SensitiveRegex, compiled.String())
				}
			}
		}
	}

	return cfg
}
