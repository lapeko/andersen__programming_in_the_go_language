package app

import (
	"fmt"
	"strings"
)

func fillWithZeros(srcString string, totalSize int) string {
	paddingSize := totalSize - len(srcString)
	return fmt.Sprintf("%s%s", strings.Repeat("0", paddingSize), srcString)
}
