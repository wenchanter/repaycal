package pmt

import (
	"context"
	"errors"
	"repaycal/apps/calculator/model/pmtlog"
	"repaycal/apps/calculator/rpc/protoc/calculator"

	"github.com/shopspring/decimal"
	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/zap"
)

// Service struct (typically in service.go)
type PmtLogService struct {
	pmtLogModel pmtlog.PmtLogModel
}

func NewPmtLogService(pmtLogModel pmtlog.PmtLogModel) *PmtLogService {
	return &PmtLogService{
		pmtLogModel: pmtLogModel,
	}
}

// calculateMonthlyRepayment computes PMT using shopspring/decimal
// Formula: PMT = p * [r(1+r)^n] / [(1+r)^n - 1]
func calculateMonthlyRepayment(loanAmount int64, interestRate int32, numberOfPayments int32) (int64, error) {
	p := decimal.NewFromInt(loanAmount)
	n := decimal.NewFromInt32(numberOfPayments)

	// Convert annual interest rate to monthly decimal: rate / 100 / 12
	r := decimal.NewFromInt32(interestRate).
		Div(decimal.NewFromInt(100)).
		Div(decimal.NewFromInt(12))

	// Edge case: zero interest rate → PMT = p / n
	if r.IsZero() {
		if n.IsZero() {
			return decimal.Zero.IntPart(), errors.New("number of payments cannot be zero")
		}
		return p.Div(n).Round(0).IntPart(), nil
	}

	// (1 + r)^n
	onePlusR := decimal.NewFromInt(1).Add(r)
	onePlusRPowN := onePlusR.Pow(n)

	// PMT = P * [r * (1+r)^n] / [(1+r)^n - 1]
	numerator := p.Mul(r.Mul(onePlusRPowN))
	denominator := onePlusRPowN.Sub(decimal.NewFromInt(1))

	if denominator.IsZero() {
		return decimal.Zero.IntPart(), errors.New("invalid inputs: denominator is zero")
	}

	pmt := numerator.Div(denominator).Round(0)
	return pmt.IntPart(), nil
}

func (s *PmtLogService) CreatePmtLog(ctx context.Context, req *calculator.PMTRequest) (*calculator.PMTResponse, error) {
	// Validate input
	if req.LoanAmount <= 0 {
		return nil, errors.New("loan amount must be greater than 0")
	}
	if req.InterestRate <= 0 {
		return nil, errors.New("interest rate must be greater than 0")
	}
	if req.NumberOfPayments <= 0 {
		return nil, errors.New("number of payments must be greater than 0")
	}

	monthlyPayment, err := calculateMonthlyRepayment(req.LoanAmount, req.InterestRate, req.NumberOfPayments)
	if err != nil {
		logx.Errorf("failed to calculate monthly repayment: %v", err)
		return nil, err
	}

	// Build the model
	data := &pmtlog.PmtLog{
		LoanAmount:       req.LoanAmount,
		InterestRate:     int64(req.InterestRate),
		NumberOfPayments: int64(req.NumberOfPayments),
		MonthlyRepayment: monthlyPayment,
	}

	// Call the Insert method
	result, err := s.pmtLogModel.Insert(ctx, data)
	if err != nil {
		logx.Errorf("failed to insert pmt log: %v", err)
		return nil, err
	}

	// Get the rows affected
	ct, err := result.RowsAffected()
	if err != nil {
		logx.Errorf("failed to get rows affected: %v", err)
		return nil, err
	}
	logx.Info("inserted pmt log", zap.Int64("rows affected ", ct))

	return &calculator.PMTResponse{
		MonthlyRepayment: monthlyPayment,
	}, nil
}
