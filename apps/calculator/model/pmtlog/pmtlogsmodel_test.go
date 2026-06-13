package pmtlog

import (
	"context"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func TestPMTlogModel_RealInsert(t *testing.T) {

	dsn := "postgres://postgres:pass123@localhost:5432/loandb?sslmode=disable"

	conn := sqlx.NewSqlConn("postgres", dsn)

	m := NewPmtLogModel(conn)

	testData := &PmtLog{
		LoanAmount:       300000,
		InterestRate:     5,
		NumberOfPayments: 24,
		MonthlyRepayment: 8774,
	}

	ctx := context.Background()

	res, err := m.Insert(ctx, testData)

	assert.NoError(t, err, "insert error")
	assert.NotNil(t, res)

	var actualCount int64
	checkQuery := `SELECT count(*) FROM "public"."pmt_log" WHERE loan_amount = $1 AND monthly_repayment = $2`
	err = conn.QueryRowCtx(ctx, &actualCount, checkQuery, testData.LoanAmount, testData.MonthlyRepayment)

	assert.NoError(t, err)

	assert.GreaterOrEqual(t, actualCount, int64(1), "at least on record")

	t.Cleanup(func() {
		cleanQuery := `DELETE FROM "public"."pmt_log" WHERE loan_amount = $1 AND monthly_repayment = $2`
		_, _ = conn.ExecCtx(ctx, cleanQuery, testData.LoanAmount)
	})
}
