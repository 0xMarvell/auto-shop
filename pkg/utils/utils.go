package utils

import (
	"fmt"
)

// NewLine prints number of new lines based on argument passed.
func NewLine(n int) {
	for ; n > 0; n-- {
		fmt.Println()
	}
}
