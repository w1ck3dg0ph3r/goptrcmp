package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/w1ck3dg0ph3r/goptrcmp"
)

func main() {
	singlechecker.Main(goptrcmp.Analyzer())
}
