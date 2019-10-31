package str_test

import (
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	str := "a.b.c"
	ret := strings.ReplaceAll(str, ".", "/")
	t.Log(ret)
}
