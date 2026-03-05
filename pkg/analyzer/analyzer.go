package analyzer

import (
	"go/ast"
	"go/token"

	"github.com/Vladroon22/linters-check/config"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var CustomAnalyzer = &analysis.Analyzer{
	Name:     "customlinter",
	Doc:      "checks log messages against predefined rules: low case words, special symbols, sensitive data and certain language",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (any, error) {
	cfg := config.NewConfig()

	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			node, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			if len(node.Args) == 0 || !IsLog(node) {
				return true
			}

			arg := node.Args[0]

			if str, ok := arg.(*ast.BasicLit); ok && str.Kind == token.STRING {
				call := str.Value[1 : len(str.Value)-1]

				if CheckLowerCase(cfg, call) {
					pass.Reportf(node.Pos(), "log message should start with a lowercase letter")
				}

				if CheckEnglishOnly(cfg, call) {
					pass.Reportf(node.Pos(), "log message should has only english symbols")
				}

				if CheckSensitiveData(cfg, call) {
					pass.Reportf(node.Pos(), "log message may contain sensitive data")
				}

				if CheckSpecialChars(cfg, call) {
					pass.Reportf(node.Pos(), "log message should not contain special characters")
				}

				if CheckEmoji(cfg, call) {
					pass.Reportf(node.Pos(), "log message should not contain emojis")
				}
			}
			return true
		})
	}

	return nil, nil
}
