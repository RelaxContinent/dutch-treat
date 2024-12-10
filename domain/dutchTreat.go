package domain

import (
	"go-capabilities-showcase/application/presenter"
	"go-capabilities-showcase/domain/types"
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
	members := req.Members
	payments := req.Payments

	// 各自の支払額マップを初期化
	paymentMap := make(map[types.Name]types.Amount, 0)
	for _, member := range members {
		paymentMap[member.Name] = 0
	}

	// 各自の支払額の和を計算
	var sum int
	for _, payment := range payments {
		sum += int(payment.Amount)
		paymentMap[payment.Payer] += payment.Amount
	}

	// 割り勘した金額を計算
	dutchTreatVal := sum / len(members) // TODO 10の位で四捨五入

	// レスポンス生成
	var res presenter.Response
	if types.Amount(dutchTreatVal) > paymentMap[members[0].Name] {
		res.Payment = presenter.Payment{
			Payer:  members[0].Name,
			Amount: types.Amount(dutchTreatVal) - paymentMap[members[0].Name],
		}
	} else {
		res.Payment = presenter.Payment{
			Payer:  members[1].Name,
			Amount: types.Amount(dutchTreatVal) - paymentMap[members[1].Name],
		}
	}

	return &res, nil
}
