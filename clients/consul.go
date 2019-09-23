package clients

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

func NewConsulClient() {

	config := api.DefaultConfig()
	fmt.Println(config)

}
