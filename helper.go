package goutil

import (
	"fmt"
	"log"
)

// FailOnError panic msg and err if error occurs
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// Must panic err if error occurs
func Must(err error) {
	if err != nil {
		log.Fatalf("%s", err)
		panic(err)
	}
}
