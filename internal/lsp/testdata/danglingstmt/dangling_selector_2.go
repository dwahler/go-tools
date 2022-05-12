package danglingstmt

import "github.com/dwahler/go-tools/internal/lsp/foo"

func _() {
	foo. //@rank(" //", Foo)
	var _ = []string{foo.} //@rank("}", Foo)
}
