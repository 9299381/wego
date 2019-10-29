package web_controller

import "github.com/9299381/wego/contracts"

type PageOne struct {
}

func (s *PageOne) Handle(ctx contracts.Context) (interface{}, error) {

	return nil, nil
}
func (s *PageOne) GetRules() interface{} {
	return nil
}
