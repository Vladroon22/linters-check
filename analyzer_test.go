package linters

import (
	"testing"

	"github.com/Vladroon22/linters-check/pkg/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, analyzer.CustomAnalyzer, "foo")
}
