package jwt

import "github.com/xyield/xumm-go-client/xumm/payload"

var transactionTypeFromString = map[string]payload.TransactionType{
	"SignIn":               payload.SignIn,
	"Payment":              payload.Payment,
	"OfferCreate":          payload.OfferCreate,
	"OfferCancel":          payload.OfferCancel,
	"EscrowFinish":         payload.EscrowFinish,
	"EscrowCreate":         payload.EscrowCreate,
	"EscrowCancel":         payload.EscrowCancel,
	"DepositPreauth":       payload.DepositPreauth,
	"CheckCreate":          payload.CheckCreate,
	"CheckCash":            payload.CheckCash,
	"CheckCancel":          payload.CheckCancel,
	"AccountSet":           payload.AccountSet,
	"PaymentChannelCreate": payload.PaymentChannelCreate,
	"PaymentChannelFund":   payload.PaymentChannelFund,
	"SetRegularKey":        payload.SetRegularKey,
	"SignerListSet":        payload.SignerListSet,
	"TrustSet":             payload.TrustSet,
	"EnableAmendment":      payload.EnableAmendment,
	"AccountDelete":        payload.AccountDelete,
	"SetFee":               payload.SetFee,
}
