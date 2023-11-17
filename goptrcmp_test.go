package goptrcmp_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/w1ck3dg0ph3r/goptrcmp"
)

func TestGoptrcmp(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, goptrcmp.Analyzer(), "./...")
}
