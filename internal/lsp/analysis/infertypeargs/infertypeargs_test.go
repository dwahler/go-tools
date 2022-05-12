// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package infertypeargs_test

import (
	"testing"

	"github.com/dwahler/go-tools/go/analysis/analysistest"
	"github.com/dwahler/go-tools/internal/lsp/analysis/infertypeargs"
	"github.com/dwahler/go-tools/internal/testenv"
	"github.com/dwahler/go-tools/internal/typeparams"
)

func Test(t *testing.T) {
	testenv.NeedsGo1Point(t, 13)
	if !typeparams.Enabled {
		t.Skip("type params are not enabled")
	}
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, infertypeargs.Analyzer, "a")
}
