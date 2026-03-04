package main

import (
	"github.com/Vladroon22/linters-check/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.CustomAnalyzer)
}
