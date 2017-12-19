package goutil

import (
	"fmt"
	"log"
)

// FailOnError panic msg when error occurs
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
