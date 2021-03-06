package mixpaysdk

/**
paymentAssetId: *required String, assetId of payment cryptocurrency.
settlementAssetId *required String, assetId of settlement cryptocurrency. Settlement assets you prefer.
payeeId *required String, three settlement modes are supported, normal user, robot, and multisig group, so it is usually the Mixin UUID of a normal user or robot, and you can also specify the multisigId of a sub-account.
quoteAmount *required Corresponding to the amount of quoteAssetId, for example, the current commodity value is 10 USD
quoteAssetId *required String, assetId of quote cryptocurrency, the asset include cryptocurrency and fiat currency.
traceId	String, UUID, varchar(36), used to prevent double payment.
clientId	String, UUID of client of the payment.
paymentAmount	Numeric, The quoteAmount parameter is invalid when paymentAmount is not null.
expireSeconds	Numeric, minimum 60, maximum 172800. This parameter is invalid when the quote currency and the payment currency are not the same.
settlementMethod	String, Instant settlement wallet. This parameter has two values, mixin and mixpay, the default is mixin.
remark	String, maximum 50. Payment remark viewable by the payee.
note	String, maximum 50. Payment note viewable by the payer.
settlementMemo	String, maximum 200. A memo similar to Mixin Snapshots, this parameter you can customize. This parameter only takes effect when your settlementMethod is equal to mixin.
*/
func (c *Client) GetPayments(req PaymentReq) (paymentResp *PaymentResp, err error) {
	response, err := c.client.R().SetBody(req).Post(PaymentsUri)
	if err != nil {
		return nil, err
	}
	if err := c.UnmarshalResponse(response, &paymentResp); err != nil {
		return paymentResp, err
	}
	return paymentResp, nil
}

func (c *Client) GetPaymentsEstimated(req PaymentReq) (resp *PaymentsEstimatedResp, err error) {
	queryParams := map[string]string{
		"paymentAssetId":    req.PaymentAssetId,
		"quoteAssetId":      req.QuoteAssetId,
		"settlementAssetId": req.SettlementAssetId,
	}
	if req.PaymentAmount != "" {
		queryParams["paymentAmount"] = req.PaymentAmount
	}
	if req.QuoteAmount != "" {
		queryParams["quoteAmount"] = req.QuoteAmount
	}
	response, err := c.client.R().
		SetQueryParams(queryParams).Get(PaymentsEstimatedUri)
	if err != nil {
		return nil, err
	}
	if err := c.UnmarshalResponse(response, &resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *Client) GetPaymentsResult(traceId string) (resp *PaymentsResultResp, err error) {
	response, err := c.client.R().SetQueryParam("traceId", traceId).Get(PaymentsResultUri)
	if err != nil {
		return nil, err
	}
	if err := c.UnmarshalResponse(response, &resp); err != nil {
		return resp, err
	}
	return resp, nil
}
