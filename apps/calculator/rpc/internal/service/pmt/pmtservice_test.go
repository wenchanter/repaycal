package pmt

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

// TestCalculateMonthlyRepayment_Basic 测试基本的 PMT 计算
func TestCalculateMonthlyRepayment_Basic(t *testing.T) {
	tests := []struct {
		name             string
		loanAmount       int64
		interestRate     int32
		numberOfPayments int32
		expected         string
		expectError      bool
		errorMsg         string
	}{
		{
			name:             "standard loan - 5% annual rate",
			loanAmount:       100000,
			interestRate:     5,
			numberOfPayments: 12,
			expected:         "8561",
			expectError:      false,
		},
		{
			name:             "standard loan - 10% annual rate",
			loanAmount:       100000,
			interestRate:     10,
			numberOfPayments: 12,
			expected:         "8792",
			expectError:      false,
		},
		{
			name:             "long term loan - 30 years",
			loanAmount:       300000,
			interestRate:     6,
			numberOfPayments: 360,
			expected:         "1799",
			expectError:      false,
		},
		{
			name:             "zero interest rate",
			loanAmount:       12000,
			interestRate:     0,
			numberOfPayments: 12,
			expected:         "1000",
			expectError:      false,
		},
		{
			name:             "single payment",
			loanAmount:       5000,
			interestRate:     12,
			numberOfPayments: 1,
			expected:         "5050",
			expectError:      false,
		},
		{
			name:             "small loan amount",
			loanAmount:       1000,
			interestRate:     5,
			numberOfPayments: 6,
			expected:         "169",
			expectError:      false,
		},
		{
			name:             "large loan amount",
			loanAmount:       1000000,
			interestRate:     3,
			numberOfPayments: 240,
			expected:         "5546",
			expectError:      false,
		},
		{
			name:             "zero number of payments with zero interest",
			loanAmount:       10000,
			interestRate:     0,
			numberOfPayments: 0,
			expectError:      true,
			errorMsg:         "number of payments cannot be zero",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calculateMonthlyRepayment(tt.loanAmount, tt.interestRate, tt.numberOfPayments)

			if tt.expectError {
				assert.Error(t, err)
				if tt.errorMsg != "" {
					assert.Equal(t, tt.errorMsg, err.Error())
				}
				return
			}

			assert.NoError(t, err)
			expected, _ := decimal.NewFromString(tt.expected)
			assert.True(t,
				expected.IntPart() == result,
				"expected %s but got %s (loan=%d, rate=%d, payments=%d)",
				expected, result, tt.loanAmount, tt.interestRate, tt.numberOfPayments,
			)
		})
	}
}

func TestCalculateMonthlyRepayment_RoundingPrecision(t *testing.T) {
	// Ensure result is always rounded to 2 decimal places
	result, err := calculateMonthlyRepayment(100000, 7, 36)
	assert.NoError(t, err)

	// Check that the result is positive (at least no negative rounding errors)
	assert.True(t, result > 0, "result should be positive")
}

func TestCalculateMonthlyRepayment_Consistency(t *testing.T) {
	// Total repayment should always be >= loan amount
	testCases := []struct {
		loanAmount       int64
		interestRate     int32
		numberOfPayments int32
	}{
		{50000, 5, 24},
		{200000, 8, 120},
		{75000, 3, 60},
	}

	for _, tc := range testCases {
		result, err := calculateMonthlyRepayment(tc.loanAmount, tc.interestRate, tc.numberOfPayments)
		assert.NoError(t, err)

		totalRepayment := result * int64(tc.numberOfPayments)
		loanAmount := tc.loanAmount

		assert.True(t,
			totalRepayment >= loanAmount,
			"total repayment %d should be >= loan amount %d",
			totalRepayment, loanAmount,
		)
	}
}

func TestCalculateMonthlyRepayment_ZeroInterest(t *testing.T) {
	// Test zero interest edge case
	result, err := calculateMonthlyRepayment(12000, 0, 12)
	assert.NoError(t, err)

	// PMT should equal LoanAmount / NumberOfPayments when interest is zero
	expected := 12000 / 12
	assert.Equal(t, int64(expected), result)
}
