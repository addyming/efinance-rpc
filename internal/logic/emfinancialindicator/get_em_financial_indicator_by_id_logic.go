package emfinancialindicator

import (
	"context"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-efinance-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-efinance-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-efinance-rpc/types/efinance"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmFinancialIndicatorByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEmFinancialIndicatorByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmFinancialIndicatorByIdLogic {
	return &GetEmFinancialIndicatorByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetEmFinancialIndicatorByIdLogic) GetEmFinancialIndicatorById(in *efinance.IDInt32Req) (*efinance.EmFinancialIndicatorInfo, error) {
	result, err := l.svcCtx.DB.EmFinancialIndicator.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &efinance.EmFinancialIndicatorInfo{
		Id:                       &result.ID,
		FinancialIndicatorsKey:   &result.FinancialIndicatorsKey,
		FinancialIndicatorsValue: &result.FinancialIndicatorsValue,
		FinancialIndicatorsName:  &result.FinancialIndicatorsName,
		CreateTime:               pointy.GetUnixMilliPointer(result.CreateTime.UnixMilli()),
		Remarks:                  &result.Remarks,
		Status:                   &result.Status,
	}, nil
}
