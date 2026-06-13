package logic

import (
	"context"
	"repaycal/apps/calculator/rpc/internal/service/pmt"
	"repaycal/apps/calculator/rpc/protoc/calculator"

	"repaycal/apps/calculator/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CalculatePMTLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCalculatePMTLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CalculatePMTLogic {
	return &CalculatePMTLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CalculatePMTLogic) CalculatePMT(in *calculator.PMTRequest) (*calculator.PMTResponse, error) {

	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err := pmt.NewPmtLogService(l.svcCtx.PmtLogModel).CreatePmtLog(l.ctx, in)
	if err != nil {
		logx.Errorf("call create pmt log err: %v", err)
		return &calculator.PMTResponse{}, err
	}
	return resp, nil
}
