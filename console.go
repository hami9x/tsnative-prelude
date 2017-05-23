package tsn

import (
	"fmt"
)

var (
	Console consoleAPI
)

type consoleAPI struct{}

func (c consoleAPI) Log(args ...interface{}) {
	for _, arg := range args {
		fmt.Print(Tsn__.ToString(arg))
		fmt.Print(" ")
	}
	fmt.Println()
}
