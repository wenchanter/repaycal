package calculator

import (
	"fmt"
)

func (x *PMTRequest) Validate() error {
	if x == nil {
		return fmt.Errorf("request body cannot be nil")
	}

	if x.LoanAmount <= 0 {
		return fmt.Errorf("loan_amount must be greater than 0, got: %d", x.LoanAmount)
	}

	if x.InterestRate <= 0 {
		return fmt.Errorf("interest_rate must be greater than 0, got: %d", x.InterestRate)
	}

	if x.NumberOfPayments <= 0 {
		return fmt.Errorf("number_of_payments must be greater than 0, got: %d", x.NumberOfPayments)
	}

	return nil
}
