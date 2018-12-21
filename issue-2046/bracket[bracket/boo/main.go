package bracket

import (
	"fmt"
	"os"

	tss "github.com/kevinburke/tss/lib"
)

func Do() {
	w := tss.NewWriter(os.Stdout)
	fmt.Fprintln(w, "vim-go")
}
