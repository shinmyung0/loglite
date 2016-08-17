package loglite

import (
	"fmt"
	"testing"
)

const (
	one = 1 << iota
	two
)

func TestBasics(t *testing.T) {

	Info("Some info logs")
	Debug("Some debug logs")
}

func variadic(args ...interface{}) {
	fmt.Println(args...)
}
