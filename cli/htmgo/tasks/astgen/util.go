package astgen

import (
	"fmt"
)

func PanicF(format string, args ...interface{}) {
	panic(fmt.Sprintf(format, args...))
}
