package goptrcmp

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

func Analyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:             "goptrcmp",
		Doc:              "Reports pointer comparisons",
		URL:              "https://github.com/w1ck3dg0ph3r/goptrcmp",
		Run:              run,
		RunDespiteErrors: false,
		Requires: []*analysis.Analyzer{
			inspect.Analyzer,
		},
	}
}

func run(pass *analysis.Pass) (any, error) {
	inspector, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, nil
	}

	filter := []ast.Node{
		(*ast.BinaryExpr)(nil),
	}

	inspector.Preorder(filter, func(node ast.Node) {
		expr, ok := node.(*ast.BinaryExpr)
		if !ok {
			return
		}

		if expr.Op != token.EQL && expr.Op != token.NEQ {
			return
		}

		analyzeComparison(pass, expr)
	})

	return nil, nil
}

func analyzeComparison(pass *analysis.Pass, expr *ast.BinaryExpr) {
	Xtype := pass.TypesInfo.TypeOf(expr.X)
	Ytype := pass.TypesInfo.TypeOf(expr.Y)
	if Xtype == nil || Ytype == nil {
		return
	}

	_, Xpointer := Xtype.(*types.Pointer)
	_, Ypointer := Ytype.(*types.Pointer)

	if Xpointer && Ypointer {
		buf := &bytes.Buffer{}
		printer.Fprint(buf, pass.Fset, expr)
		pass.Reportf(expr.Pos(), "pointer comparison: %s", buf.Bytes())
	}
}
