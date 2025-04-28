package domain

import (
	"dutch-treat/application/presenter"
	"dutch-treat/domain/types"
)

type DutchTreatService interface {
	DutchTreat(*presenter.Request) (*presenter.Response, error)
}

type dutchTreatService struct {
}

func NewDutchTreatService() DutchTreatService {
	return &dutchTreatService{}
}

func (s dutchTreatService) DutchTreat(req *presenter.Request) (*presenter.Response, error) {
	// 各人の支払いを集計する
	paymentMap := make(map[types.Name]types.Amount)
	for _, payment := range req.Payments {
		paymentMap[payment.Payer] += payment.Amount
	}

	// 支払いを合計
	var allPayment types.Amount
	for _, payment := range paymentMap {
		allPayment += payment
	}

	// 割り勘の金額を計算する
	dutchTreatAmount := allPayment / types.Amount(len(paymentMap))

	// 各人の支払いを計算する
	result := make([]presenter.Payment, len(paymentMap))
	i := 0
	for name, amount := range paymentMap {
		result[i] = presenter.Payment{
			Payer:  name,
			Amount: dutchTreatAmount - amount,
		}
		i++
	}

	return &presenter.Response{Payment: result}, nil
}
