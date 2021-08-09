
import (
	"strings"
)

func generateTheString(n int) string {
	if n%2 == 0 {
		return strings.Repeat("b", n-1) + "a"
	} else {
		return strings.Repeat("a", n)
	}
}
