module github.com/9299381/wego/demo

go 1.12

require (
	github.com/9299381/wego v0.1.0
	github.com/Azure/go-autorest v10.15.3+incompatible
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/hashicorp/consul/api v1.2.0
	github.com/looplab/fsm v0.1.0
	github.com/mitchellh/mapstructure v1.1.2
	github.com/tidwall/gjson v1.3.4
	go.mongodb.org/mongo-driver v1.1.2
	xorm.io/builder v0.3.5
)

replace github.com/9299381/wego => ../
