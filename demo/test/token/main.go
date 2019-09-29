package main

import (
	"fmt"
	"github.com/9299381/wego/tools/jwt"
)

func main() {
	token := jwt.New().
		SetId("123").
		SetName("abc").
		SetRole("p1101").
		GetToken()
	//eyJpZCI6IjEyMyIsIm5hbWUiOiJhYmMiLCJyb2xlIjoicDExMDEiLCJpYXQiOjE1Njk1NzkwOTQsImV4cCI6MTU3MjE3MTA5NH0=.228a1c01bdce8cd3405d2295d7a2eb30
	fmt.Println(token)
}
