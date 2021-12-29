package common

import (
	"fmt"
)

func Log(message interface{}) {
	fmt.Printf("%#v\n", message)
}
