package svc

import (
	"repaycal/apps/calculator/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"repaycal/apps/calculator/model/pmtlog"

	_ "github.com/lib/pq"
)

type ServiceContext struct {
	Config      config.Config
	PmtLogModel pmtlog.PmtLogModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn("postgres", c.Postgresql.DataSource)

	return &ServiceContext{
		Config:      c,
		PmtLogModel: pmtlog.NewPmtLogModel(conn),
	}
}
