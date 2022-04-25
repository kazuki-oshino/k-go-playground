package secretstring

type ConfidentialInfo struct {
	Id             int64
	CreditCardCode CreditCardCode
}

type CreditCardCode int

func NewConfidentialInfo(id int64, creditCardCode CreditCardCode) *ConfidentialInfo {
	return &ConfidentialInfo{
		Id:             id,
		CreditCardCode: creditCardCode,
	}
}

func (c CreditCardCode) String() string {
	return "xxxx-xxxx-xxxx-xxxx"
}

func (c CreditCardCode) GoString() string {
	return "xxxx-xxxx-xxxx-xxxx"
}
