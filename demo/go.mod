module github.com/9299381/wego/demo

go 1.12

require (
	github.com/9299381/wego v0.2.0
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/looplab/fsm v0.1.0
	github.com/mitchellh/mapstructure v1.1.2
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/shopspring/decimal v0.0.0-20180709203117-cd690d0c9e24
	github.com/tidwall/gjson v1.3.4
	golang.org/x/net v0.0.0-20190603091049-60506f45cf65 // indirect
	golang.org/x/tools v0.0.0-20190606050223-4d9ae51c2468 // indirect
	gopkg.in/mgo.v2 v2.0.0-20160818020120-3f83fa500528
	xorm.io/builder v0.3.7
)

replace github.com/9299381/wego => ../
