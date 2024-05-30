package services

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"safpass-api/configs"
)

type MidtransService struct {
	snap snap.Client
}

func NewMidtransService() *MidtransService {
	config := configs.LoadConfig()
	s := snap.Client{}
	s.New(config.MidtransServerKey, midtrans.Sandbox)

	return &MidtransService{
		snap: s,
	}
}

// CreateTransaction creates a new transaction in Midtrans
func (m *MidtransService) CreateTransaction(orderID string, grossAmount int64, customerDetails *midtrans.CustomerDetails) (*snap.Response, error) {
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: grossAmount,
		},
		Gopay: &snap.GopayDetails{
			EnableCallback: true,
		},
		CustomerDetail: customerDetails,
	}

	snapResp, err := m.snap.CreateTransaction(req)
	if err != nil {
		return nil, err
	}
	return snapResp, nil
}
