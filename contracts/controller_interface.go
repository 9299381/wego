package contracts

type IController interface {
	Handle(ctx Context) (interface{}, error)
}

type Controller struct {
}

func (s *Controller) Handle(ctx Context) (interface{}, error) {
	return nil, nil
}
func (s *Controller) GetRules() interface{} {
	return nil
}
func (s *Controller) Mock() interface{} {
	return nil
}
