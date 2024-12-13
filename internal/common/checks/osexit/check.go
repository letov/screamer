package osexit

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name:             "os_exit_check",
	Doc:              "os_exit_check",
	Run:              run,
	RunDespiteErrors: false,
}

const errorMessage = "os.Exit usage"

func isIdent(n ast.Expr, names ...string) bool {
	switch x := n.(type) {
	case *ast.Ident:
		for _, n := range names {
			if x.Name == n {
				return true
			}
		}
	}
	return false
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		ast.Inspect(f, func(n ast.Node) bool {
			switch x := n.(type) {
			case *ast.SelectorExpr:
				if isIdent(x.X, "os") && x.Sel.Name == "Exit" {
					pass.Reportf(n.Pos(), errorMessage)
					return false
				}
			}
			return true
		})
	}
	return nil, nil
}
