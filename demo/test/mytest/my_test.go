package mytest

import (
	"fmt"
	"testing"
)

func init() {
	testing.Init()
}
func TestSomething(t *testing.T) {
	fmt.Println("here")
}
