package main

import (
	"github.com/24h-purewater/mixpay-sdk-go"
	"log"
)

const (
	BTC  = "c6d0c728-2624-429b-8e0d-d9d19b6592fa"
	USDT = "4d8c508b-91c5-375b-92b0-ee702ed2dac5"
	USDC = "9b180ab6-6abe-3dc0-a13f-04169eb34bfa"
)

func main() {
	mixpay := mixpaysdkgo.New()
	//eg. use btc to swap usdc
	req := mixpaysdkgo.PaymentReq{
		PaymentAssetId:    BTC,
		PaymentAmount:     "0.01",
		SettlementAssetId: USDC,
		PayeeId:           "019308a5-e1a9-427c-af2a-e05093beedaa",
		QuoteAssetId:      USDC,
	}
	paymentsResp, err := mixpay.GetPayments(req)
	if err != nil {
		log.Fatalln(err)
	}
	if paymentsResp.Code != 0 {
		log.Printf("get payments error: %+v", paymentsResp)
	}
	log.Printf("mixin payments: %+v", paymentsResp.Data)
}