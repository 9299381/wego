package idwork

import (
	"github.com/9299381/wego/configs"
	"strconv"
)

func ID() string {
	serverId, _ := strconv.Atoi(configs.Env("SERVER_ID", "512"))
	return getID(int64(serverId))
}
