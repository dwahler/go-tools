package other

import "github.com/dwahler/go-tools/internal/lsp/rename/crosspkg"

func Other() {
	crosspkg.Bar
	crosspkg.Foo() //@rename("Foo", "Flamingo")
}
