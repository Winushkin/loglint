package analyzer

import (
	"go/ast"

	"strings"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "loglint",
	Doc:  "Checks logs messages",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}
			if !isLoggerCall(call) {
				return true
			}
			checkCall(pass, call)
			return true
		})
	}
	return nil, nil
}

func isLoggerCall(call *ast.CallExpr) bool {
	fun, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	method := fun.Sel.Name

	switch method {
	case "Info", "Error", "Warn", "Debug":
		return true
	}
	return false
}

func checkCall(pass *analysis.Pass, call *ast.CallExpr) {
	if len(call.Args) == 0 {
		return
	}
	arg, ok := call.Args[0].(*ast.BasicLit)
	if !ok {
		return
	}

	if arg.Kind.String() != "STRING" {
		return
	}
	msg := strings.Trim(arg.Value, `"`)

	checkLowerStart(pass, call, msg)
	checkEnglish(pass, call, msg)
	checkSpecialChars(pass, call, msg)
	checkSensitiveData(pass, call, msg)
}
