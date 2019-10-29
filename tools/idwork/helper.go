package idwork

import (
	"github.com/9299381/wego/configs"
)

func ID() string {
	return getID(int64(configs.EnvInt("server_id", 512)))
}
