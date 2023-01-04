package mixpaysdk

type (
	MixPayResp[T any] struct {
		Code        int64  `json:"code"`
		Success     bool   `json:"success"`
		Message     string `json:"message"`
		Data        T      `json:"data"`
		TimestampMS int64  `json:"timestampMs"`
	}

	Payment struct {
		IsChain                   bool   `json:"isChain"`
		Expire                    int64  `json:"expire"`
		Seconds                   int64  `json:"seconds"`
		PayeeID                   string `json:"payeeId"`
		TraceID                   string `json:"traceId"`
		ClientID                  string `json:"clientId"`
		PaymentAssetID            string `json:"paymentAssetId"`
		SettlementAssetID         string `json:"settlementAssetId"`
		QuoteAssetID              string `json:"quoteAssetId"`
		PaymentAmount             string `json:"paymentAmount"`
		QuoteAmount               string `json:"quoteAmount"`
		EstimatedSettlementAmount string `json:"estimatedSettlementAmount"`
		Memo                      string `json:"memo"`
		Recipient                 string `json:"recipient"`
		SettlementAssetSymbol     string `json:"settlementAssetSymbol"`
		PaymentAssetSymbol        string `json:"paymentAssetSymbol"`
		QuoteAssetSymbol          string `json:"quoteAssetSymbol"`
		Destination               string `json:"destination"`
		Tag                       string `json:"tag"`
	}

	PaymentResp = MixPayResp[Payment]

	PaymentReq struct {
		PaymentAssetId    string `json:"paymentAssetId,omitempty"`
		SettlementAssetId string `json:"settlementAssetId,omitempty"`
		PayeeId           string `json:"payeeId,omitempty"`
		QuoteAmount       string `json:"quoteAmount,omitempty"`
		QuoteAssetId      string `json:"quoteAssetId,omitempty"`
		TraceId           string `json:"traceId,omitempty"`
		ClientId          string `json:"clientId,omitempty"`
		PaymentAmount     string `json:"paymentAmount,omitempty"`
		ExpireSeconds     string `json:"expireSeconds,omitempty"`
		SettlementMethod  string `json:"settlementMethod,omitempty"`
		Remark            string `json:"remark,omitempty"`
		Note              string `json:"note,omitempty"`
		SettlementMemo    string `json:"settlementMemo,omitempty"`
	}

	PaymentsEstimated struct {
		Price                     string `json:"price"`
		EstimatedSettlementAmount string `json:"estimatedSettlementAmount"`
		SettlementAssetID         string `json:"settlementAssetId"`
		SettlementAssetSymbol     string `json:"settlementAssetSymbol"`
		PaymentAssetID            string `json:"paymentAssetId"`
		PaymentAssetSymbol        string `json:"paymentAssetSymbol"`
		PaymentAmount             string `json:"paymentAmount"`
		QuoteAssetSymbol          string `json:"quoteAssetSymbol"`
		QuoteAssetID              string `json:"quoteAssetId"`
		QuoteAmount               string `json:"quoteAmount"`
	}

	PaymentsEstimatedResp = MixPayResp[PaymentsEstimated]

	PaymentsResult struct {
		Status           string      `json:"status"`
		QuoteAmount      string      `json:"quoteAmount"`
		QuoteSymbol      string      `json:"quoteSymbol"`
		QuoteAssetID     string      `json:"quoteAssetId"`
		PaymentAmount    string      `json:"paymentAmount"`
		PaymentSymbol    string      `json:"paymentSymbol"`
		PaymentAssetID   string      `json:"paymentAssetId"`
		Payee            string      `json:"payee"`
		PayeeID          string      `json:"payeeId"`
		PayeeMixinNumber string      `json:"payeeMixinNumber"`
		PayeeAvatarURL   string      `json:"payeeAvatarUrl"`
		Txid             string      `json:"txid"`
		BlockExplorerURL string      `json:"blockExplorerUrl"`
		Date             any         `json:"date"`
		SurplusAmount    string      `json:"surplusAmount"`
		SurplusStatus    string      `json:"surplusStatus"`
		Confirmations    int64       `json:"confirmations"`
		PayableAmount    string      `json:"payableAmount"`
		FailureCode      string      `json:"failureCode"`
		FailureReason    string      `json:"failureReason"`
		ReturnTo         string      `json:"returnTo"`
		TraceID          string      `json:"traceId"`
		Nonce            interface{} `json:"nonce"`
		Remark           string      `json:"remark"`
	}
	PaymentsResultResp = MixPayResp[PaymentsResult]
)
