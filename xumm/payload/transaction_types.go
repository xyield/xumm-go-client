package payload

import "encoding/json"

type TransactionType uint32

type transactionType struct {
	Type TransactionType `json:"TransactionType,omitempty"`
}

type TransactionTypeError struct {
}

func (e *TransactionTypeError) Error() string {
	return "Invalid transaction type provided."
}

const (
	// Invalid TransactionType = iota
	SignIn TransactionType = 1 << iota
	Payment
	OfferCreate
	OfferCancel
	EscrowFinish
	EscrowCreate
	EscrowCancel
	DepositPreauth
	CheckCreate
	CheckCash
	CheckCancel
	AccountSet
	PaymentChannelCreate
	PaymentChannelFund
	SetRegularKey
	SignerListSet
	TrustSet
	EnableAmendment
	AccountDelete
	SetFee
)

var transactionTypeToString = map[TransactionType]string{
	SignIn:               "SignIn",
	Payment:              "Payment",
	OfferCreate:          "OfferCreate",
	OfferCancel:          "OfferCancel",
	EscrowFinish:         "EscrowFinish",
	EscrowCreate:         "EscrowCreate",
	EscrowCancel:         "EscrowCancel",
	DepositPreauth:       "DepositPreauth",
	CheckCreate:          "CheckCreate",
	CheckCash:            "CheckCash",
	CheckCancel:          "CheckCancel",
	AccountSet:           "AccountSet",
	PaymentChannelCreate: "PaymentChannelCreate",
	PaymentChannelFund:   "PaymentChannelFund",
	SetRegularKey:        "SetRegularKey",
	SignerListSet:        "SignerListSet",
	TrustSet:             "TrustSet",
	EnableAmendment:      "EnableAmendment",
	AccountDelete:        "AccountDelete",
	SetFee:               "SetFee",
}

var transactionTypeFromString = map[string]TransactionType{
	"SignIn":               SignIn,
	"Payment":              Payment,
	"OfferCreate":          OfferCreate,
	"OfferCancel":          OfferCancel,
	"EscrowFinish":         EscrowFinish,
	"EscrowCreate":         EscrowCreate,
	"EscrowCancel":         EscrowCancel,
	"DepositPreauth":       DepositPreauth,
	"CheckCreate":          CheckCreate,
	"CheckCash":            CheckCash,
	"CheckCancel":          CheckCancel,
	"AccountSet":           AccountSet,
	"PaymentChannelCreate": PaymentChannelCreate,
	"PaymentChannelFund":   PaymentChannelFund,
	"SetRegularKey":        SetRegularKey,
	"SignerListSet":        SignerListSet,
	"TrustSet":             TrustSet,
	"EnableAmendment":      EnableAmendment,
	"AccountDelete":        AccountDelete,
	"SetFee":               SetFee,
}

func (t *TransactionType) String() string {
	if s, ok := transactionTypeToString[*t]; ok {
		return s
	}
	return "Unknown"
}

func (t *TransactionType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	var v TransactionType
	var ok bool
	if v, ok = transactionTypeFromString[s]; !ok {
		return &TransactionTypeError{}
	}

	*t = v
	return nil
}
