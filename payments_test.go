package mixpaysdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	BTC  = "c6d0c728-2624-429b-8e0d-d9d19b6592fa"
	USDT = "4d8c508b-91c5-375b-92b0-ee702ed2dac5"
	USDC = "9b180ab6-6abe-3dc0-a13f-04169eb34bfa"
)

func TestGetPaymentResult(t *testing.T) {
	assert := assert.New(t)
	client := New()
	resp, err := client.GetPaymentsResult("ba00b5e6-ddba-4b14-aede-c4c5e2ea51d4")
	assert.Nil(err)
	assert.Equal("success", resp.Data.Status)

	resp, err = client.GetPaymentsResult("ae69a0c0-73b9-4698-a5bb-a1add434a680")
	assert.Nil(err)
	assert.Equal("unpaid", resp.Data.Status)
}



func TestGetPayments(t *testing.T) {
	assert := assert.New(t)
	client := New()

	req := PaymentReq{
		PaymentAssetId:    BTC,
		PaymentAmount:     "0.01",
		SettlementAssetId: USDC,
		PayeeId:           "019308a5-e1a9-427c-af2a-e05093beedaa",
		QuoteAssetId:      USDC,
	}
	paymentsResp, err := client.GetPayments(req)
	assert.Nil(err)
	assert.NotEmpty(paymentsResp.Data.EstimatedSettlementAmount)
}

func TestGetPaymentsEstimated(t *testing.T) {
	assert := assert.New(t)
	client := New()

	req := PaymentReq{
		PaymentAssetId:    BTC,
		PaymentAmount:     "0.01",
		SettlementAssetId: USDC,
		QuoteAssetId:      USDC,
	}
	resp, err := client.GetPaymentsEstimated(req)
	assert.Nil(err)
	assert.NotEmpty(resp.Data.EstimatedSettlementAmount)
}