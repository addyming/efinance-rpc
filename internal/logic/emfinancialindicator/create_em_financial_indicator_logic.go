package emfinancialindicator

import (
	"context"

	"github.com/suyuan32/simple-admin-efinance-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-efinance-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-efinance-rpc/types/efinance"

    "github.com/suyuan32/simple-admin-common/msg/errormsg"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateEmFinancialIndicatorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateEmFinancialIndicatorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateEmFinancialIndicatorLogic {
	return &CreateEmFinancialIndicatorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateEmFinancialIndicatorLogic) CreateEmFinancialIndicator(in *efinance.EmFinancialIndicatorInfo) (*efinance.BaseIDInt32Resp, error) {
    result, err := l.svcCtx.DB.EmFinancialIndicator.Create().
			SetNotNilFinancialIndicatorsKey(in.FinancialIndicatorsKey).
			SetNotNilFinancialIndicatorsValue(in.FinancialIndicatorsValue).
			SetNotNilFinancialIndicatorsName(in.FinancialIndicatorsName).
			SetNotNilCreateTime(pointy.GetTimeMilliPointer(in.CreateTime)).
			SetNotNilRemarks(in.Remarks).
			SetNotNilStatus(in.Status).
			Save(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &efinance.BaseIDInt32Resp{Id: result.ID, Msg: errormsg.CreateSuccess }, nil
}
