package domain

type OperationTypeID int64

const (
	Purchase               OperationTypeID = 1
	PurchaseInInstallments                 = 2
	Withdraw                               = 3
	Payment                                = 4
)

func (o OperationTypeID) IsNegative() bool {
	switch o {
	case Purchase, PurchaseInInstallments:
		return true
	case Withdraw:
		return true
	case Payment:
		return false
	default:
		return true
	}
}

func (o OperationTypeID) Validate(amount int64) error {
	if o.IsNegative() && amount > 0 {
		return ErrorAmountMustBeNegativeForThatOperationType
	}
	if !o.IsNegative() && amount < 0 {
		return ErrorAmountMustBePositiveForThatOperationType
	}
	return nil
}

type OperationType struct {
	ID            int64  `json:"id"`
	DescriptionPT string `json:"description_pt"`
}
