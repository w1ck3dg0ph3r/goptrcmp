package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"os"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/analysis/singlechecker"
	"golang.org/x/tools/go/ast/inspector"
)

func main() {
	singlechecker.Main(New())
}

func New() *analysis.Analyzer {
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

	var runerr error

	inspector.Preorder(filter, func(node ast.Node) {
		if runerr != nil {
			return
		}

		n, ok := node.(*ast.BinaryExpr)
		if !ok {
			return
		}

		if n.Op != token.EQL && n.Op != token.NEQ {
			return
		}

		if err := analyzeComparison(pass, n); err != nil {
			runerr = err
			return
		}
	})

	return nil, runerr
}

var sources = map[string][]byte{}

func analyzeComparison(pass *analysis.Pass, node *ast.BinaryExpr) error {
	Xtype, ok := pass.TypesInfo.Types[node.X]
	if !ok {
		return nil
	}

	Ytype, ok := pass.TypesInfo.Types[node.Y]
	if !ok {
		return nil
	}

	_, Xpointer := Xtype.Type.(*types.Pointer)
	_, Ypointer := Ytype.Type.(*types.Pointer)

	if Xpointer && Ypointer {
		pos := pass.Fset.Position(node.Pos())
		src, ok := sources[pos.Filename]
		if !ok {
			var err error
			src, err = os.ReadFile(pos.Filename)
			if err != nil {
				return fmt.Errorf("read source file: %w", err)
			}
			sources[pos.Filename] = src
		}

		left := src[pass.Fset.Position(node.X.Pos()).Offset:pass.Fset.Position(node.X.End()).Offset]
		right := src[pass.Fset.Position(node.Y.Pos()).Offset:pass.Fset.Position(node.Y.End()).Offset]

		op := "=="
		if node.Op == token.NEQ {
			op = "!="
		}

		pass.Reportf(node.Pos(), "pointer comparison: %s %s %s", left, op, right)
	}

	return nil
}
