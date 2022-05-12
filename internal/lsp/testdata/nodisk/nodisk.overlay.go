package nodisk

import (
	"github.com/dwahler/go-tools/internal/lsp/foo"
)

func _() {
	foo.Foo() //@complete("F", Foo, IntFoo, StructFoo)
}
