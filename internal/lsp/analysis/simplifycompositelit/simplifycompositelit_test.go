// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package simplifycompositelit_test

import (
	"testing"

	"github.com/dwahler/go-tools/go/analysis/analysistest"
	"github.com/dwahler/go-tools/internal/lsp/analysis/simplifycompositelit"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, simplifycompositelit.Analyzer, "a")
}
