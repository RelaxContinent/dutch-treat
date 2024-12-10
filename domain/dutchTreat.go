package domain

type DutchTreatService interface {
	DutchTreat() error
}

type dutchTreatService struct {
}

func NewDutchTreatService() DutchTreatService {
	return &dutchTreatService{}
}

func (s dutchTreatService) DutchTreat() error {
	return nil
}
