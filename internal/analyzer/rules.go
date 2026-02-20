package analyzer

import (
	"go/ast"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

var sensitiveKeywords = []string{
	"password",
	"token",
	"key",
	"secret",
}

func checkLowerStart(pass *analysis.Pass, call *ast.CallExpr, msg string){
	if len(msg) == 0{
		return
	}

	for _, r := range msg{
		if unicode.IsUpper(r){
			pass.Reportf(call.Pos(), "log message must start with lowercase letter")
		}
		break
	}
}

func checkEnglish(pass *analysis.Pass, call *ast.CallExpr, msg string){
	for _, r := range msg{
		if r > unicode.MaxASCII{
			pass.Reportf(call.Pos(), "log message must be in English only")
			return
		}
	}
}	


func checkSpecialChars(pass *analysis.Pass, call *ast.CallExpr, msg string){
	for _, r := range msg{
		if !(unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r)){
			pass.Reportf(call.Pos(), "log message contains special characters")
			return
		}
	}
}

func checkSensitiveData(pass *analysis.Pass, call *ast.CallExpr, msg string){
	lower := strings.ToLower(msg)

	for _, kw := range sensitiveKeywords{
		if strings.Contains(lower, kw){
			pass.Reportf(call.Pos(), "log message may contain sensitive data")
			return
		}
	}

	for _, arg := range call.Args{
		v, ok := arg.(*ast.Ident)
		if ok{
			for _, kw := range sensitiveKeywords{
				varName := strings.ToLower(v.Name)
				if strings.Contains(varName, kw){
					pass.Reportf(call.Pos(), "log message may contain sensitive data")
					return
				}
			}
		}
	}
}