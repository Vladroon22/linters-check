// plugin/main.go
package main

import (
	"github.com/Vladroon22/linters-check/analyzer"
	"golang.org/x/tools/go/analysis"
)

func New() []*analysis.Analyzer {
	return []*analysis.Analyzer{analyzer.CustomAnalyzer}
}
