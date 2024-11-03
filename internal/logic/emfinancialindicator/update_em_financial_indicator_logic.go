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

type UpdateEmFinancialIndicatorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateEmFinancialIndicatorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateEmFinancialIndicatorLogic {
	return &UpdateEmFinancialIndicatorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateEmFinancialIndicatorLogic) UpdateEmFinancialIndicator(in *efinance.EmFinancialIndicatorInfo) (*efinance.BaseResp, error) {
	err:= l.svcCtx.DB.EmFinancialIndicator.UpdateOneID(*in.Id).
			SetNotNilFinancialIndicatorsKey(in.FinancialIndicatorsKey).
			SetNotNilFinancialIndicatorsValue(in.FinancialIndicatorsValue).
			SetNotNilFinancialIndicatorsName(in.FinancialIndicatorsName).
			SetNotNilCreateTime(pointy.GetTimeMilliPointer(in.CreateTime)).
			SetNotNilRemarks(in.Remarks).
			SetNotNilStatus(in.Status).
			Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &efinance.BaseResp{Msg: errormsg.UpdateSuccess }, nil
}
