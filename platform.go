package platform

import (
	"fmt"
	"os"
)

func init() {
	// Safe PoC: prints to stderr when imported
	fmt.Fprintln(os.Stderr, "[POC] hijacked module executed")
}
