package services

import (
	"github.com/9299381/wego/contracts"
)

type CommService struct {
}

func (it *CommService) Next(srv contracts.IService) contracts.IService {
	return it
}

func (it *CommService) Handle(ctx contracts.Context) error {
	return nil
}
