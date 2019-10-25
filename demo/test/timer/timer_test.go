package test

import (
	"fmt"
	"github.com/9299381/wego/tools"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	now := tools.LocalTime(time.Now())
	n := fmt.Sprintf("%s", now)
	fmt.Println(n)
}
