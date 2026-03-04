package analyzer

import (
	"go/ast"
	"go/token"

	"github.com/Vladroon22/linters-check/config"
	"golang.org/x/tools/go/analysis"
)

var AnalyzerPlugin = map[string]*analysis.Analyzer{
	"customlinter": CustomAnalyzer,
}

var CustomAnalyzer = &analysis.Analyzer{
	Name: "customlinter",
	Doc:  "checks log messages against predefined rules",
	Run:  run,
}

func run(pass *analysis.Pass) (any, error) {
	cfg := config.NewConfig()

	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			node, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			if len(node.Args) == 0 || !isLog(node) {
				return true
			}

			arg := node.Args[0]

			if str, ok := arg.(*ast.BasicLit); ok && str.Kind == token.STRING {
				call := str.Value[1 : len(str.Value)-1]

				if !CheckLowerCase(cfg, call) {
					pass.Report(analysis.Diagnostic{
						Pos:     node.Pos(),
						End:     node.End(),
						Message: "log message should start with a lowercase letter",
					})
				}

				if !CheckEnglishOnly(cfg, call) {
					pass.Report(analysis.Diagnostic{
						Pos:     node.Pos(),
						End:     node.End(),
						Message: "log message should has only english symbols",
					})
				}

				if !CheckSensitiveData(cfg, call) {
					pass.Report(analysis.Diagnostic{
						Pos:     node.Pos(),
						End:     node.End(),
						Message: "log message may contain sensitive data",
					})
				}

				if !CheckSpecialChars(cfg, call) {
					pass.Report(analysis.Diagnostic{
						Pos:     node.Pos(),
						End:     node.End(),
						Message: "log message should not contain special characters",
					})
				}

				if !CheckEmoji(cfg, call) {
					pass.Report(analysis.Diagnostic{
						Pos:     node.Pos(),
						End:     node.End(),
						Message: "log message should not contain emojis",
					})
				}

			}
			return true
		})
	}

	return nil, nil
}

/*
func f () {


	if cfg.AutoFix {
		modifiedMsg = fixSensitiveData(cfg, modifiedMsg)
		hasIssues = true
	} else {
		// Только предлагаем исправление, но не применяем автоматически

		},
	}
	}
	pass.Report(diagnostic)
	}
}
*/
