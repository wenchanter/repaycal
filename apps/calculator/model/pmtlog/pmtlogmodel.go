package pmtlog

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PmtLogModel = (*customPmtLogModel)(nil)

type (
	// PmtLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmtLogModel.
	PmtLogModel interface {
		pmtLogModel
		withSession(session sqlx.Session) PmtLogModel
	}

	customPmtLogModel struct {
		*defaultPmtLogModel
	}
)

// NewPmtLogModel returns a model for the database table.
func NewPmtLogModel(conn sqlx.SqlConn) PmtLogModel {
	return &customPmtLogModel{
		defaultPmtLogModel: newPmtLogModel(conn),
	}
}

func (m *customPmtLogModel) withSession(session sqlx.Session) PmtLogModel {
	return NewPmtLogModel(sqlx.NewSqlConnFromSession(session))
}
