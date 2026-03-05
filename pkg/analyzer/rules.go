package analyzer

import (
	"go/ast"
	"regexp"
	"strings"
	"unicode"

	"github.com/Vladroon22/linters-check/config"
)

func IsLog(node *ast.CallExpr) bool {
	if fun, ok := node.Fun.(*ast.SelectorExpr); ok {
		switch fun.Sel.Name {
		case "Info", "Error", "Warn", "Debug", "Println", "Print", "Printf", "Fatal", "Fatalf", "Fatalln", "DPanic", "Panic":
			return true
		}
	}

	return false
}

func CheckLowerCase(cfg *config.Config, msg string) bool {
	if !cfg.CheckLowercase || msg == "" {
		return true
	}

	firstChar := []rune(msg)[0]
	if !unicode.IsUpper(firstChar) {
		return false
	}

	return true
}

func CheckEnglishOnly(cfg *config.Config, msg string) bool {
	if !cfg.CheckEnglish || msg == "" {
		return true
	}

	for _, r := range msg {
		if !unicode.Is(unicode.Latin, r) && !unicode.IsSpace(r) {
			return false
		}
	}

	return true
}

func CheckSpecialChars(cfg *config.Config, msg string) bool {
	if !cfg.CheckSpecialChars || msg == "" {
		return true
	}

	for _, r := range msg {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && !unicode.IsSpace(r) && !unicode.IsPunct(r) {
			return false
		}
	}

	return true
}

func CheckEmoji(cfg *config.Config, msg string) bool {
	if !cfg.CheckSpecialChars || msg == "" {
		return true
	}

	for _, r := range msg {
		if r > 0xFFFF || (r >= 0x1F600 && r <= 0x1F64F) || (r >= 0x1F300 && r <= 0x1F5FF) || (r >= 0x1F680 && r <= 0x1F6FF) {
			return false
		}
	}

	return true
}

func CheckSensitiveData(cfg *config.Config, msg string) bool {
	if !cfg.CheckSensitiveData || msg == "" {
		return true
	}

	lowerMsg := strings.ToLower(msg)
	for _, rgx := range cfg.SensitiveRegex {
		if ok, _ := regexp.MatchString(rgx, lowerMsg); !ok {
			return false
		}
	}

	for _, keyword := range cfg.SensitiveKeywords {
		if strings.Contains(lowerMsg, keyword) {
			return false
		}
	}

	return true
}

/*
func SuggestedFixSensitive(cfg *config.Config.config.Config, call *ast.CallExpr) (*ast.Ident, bool) {
	if !cfg.CheckSensitiveData {
		return nil, true
	}

	ident, ok := call.Fun.(*ast.Ident)
	if !ok {
		return nil, true
	}

	name := strings.ToLower(ident.Name)
	for _, key := range cfg.SensitiveKeywords {
		if strings.Contains(name, key) {
			return ident, false
		}
	}

	return nil, true
}
*/
