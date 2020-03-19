package demo

import (
	"fmt"
	"strings"
)

type Man struct {
}

func (s *Man) Work() {
	fmt.Println("man working")
}
func (s *Man) Write(content string) string {
	str := "man write something"
	ret := strings.Join([]string{str, content}, ":")
	return ret
}
