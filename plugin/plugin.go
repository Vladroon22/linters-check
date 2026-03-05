package plugin

import (
	"github.com/Vladroon22/linters-check/pkg/analyzer"
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("customlinter", New)
}

func New(conf any) (register.LinterPlugin, error) {
	return &plugin{}, nil
}

type plugin struct {
}

func (p *plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{analyzer.CustomAnalyzer}, nil
}

func (*plugin) GetLoadMode() string { return register.LoadModeTypesInfo }
