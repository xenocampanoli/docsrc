//	report2.go
package reports

import (
	"fmt"
	"strings"
)

func genLineNoPage(cStr string) (oStr string) {
	la := strings.Split(cStr, "\n")
	l := 1
	oStr = ""
	for _, s := range la {
		oStr += fmt.Sprintf("%d %s\n", l, s)
		l += 1
	}
	return
}

// End of common.go
