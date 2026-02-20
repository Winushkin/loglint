package main


import(
	"github.com/Winushkin/loglint/internal/analyzer"
	"golang.org/x/tools/go/analysis"
)


var AnalyzerPlugin = map[string]*analysis.Analyzer{
	"loglint": analyzer.Analyzer,
}