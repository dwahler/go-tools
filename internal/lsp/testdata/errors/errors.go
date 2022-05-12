package errors

import (
	"github.com/dwahler/go-tools/internal/lsp/types"
)

func _() {
	bob.Bob() //@complete(".")
	types.b //@complete(" //", Bob_interface)
}
