package analyzer

import (
	"testing"

	"github.com/Vladroon22/linters-check/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, analyzer.CustomAnalyzer, "foo")
}
