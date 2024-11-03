package emfinancialindicator

import (
	"context"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-efinance-rpc/ent/emfinancialindicator"
	"github.com/suyuan32/simple-admin-efinance-rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-efinance-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-efinance-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-efinance-rpc/types/efinance"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmFinancialIndicatorListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEmFinancialIndicatorListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmFinancialIndicatorListLogic {
	return &GetEmFinancialIndicatorListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetEmFinancialIndicatorListLogic) GetEmFinancialIndicatorList(in *efinance.EmFinancialIndicatorListReq) (*efinance.EmFinancialIndicatorListResp, error) {
	var predicates []predicate.EmFinancialIndicator
	if in.FinancialIndicatorsKey != nil {
		predicates = append(predicates, emfinancialindicator.FinancialIndicatorsKeyContains(*in.FinancialIndicatorsKey))
	}
	if in.FinancialIndicatorsValue != nil {
		predicates = append(predicates, emfinancialindicator.FinancialIndicatorsValueContains(*in.FinancialIndicatorsValue))
	}
	if in.FinancialIndicatorsName != nil {
		predicates = append(predicates, emfinancialindicator.FinancialIndicatorsNameContains(*in.FinancialIndicatorsName))
	}
	result, err := l.svcCtx.DB.EmFinancialIndicator.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &efinance.EmFinancialIndicatorListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &efinance.EmFinancialIndicatorInfo{
			Id:                       &v.ID,
			FinancialIndicatorsKey:   &v.FinancialIndicatorsKey,
			FinancialIndicatorsValue: &v.FinancialIndicatorsValue,
			FinancialIndicatorsName:  &v.FinancialIndicatorsName,
			CreateTime:               pointy.GetUnixMilliPointer(v.CreateTime.UnixMilli()),
			Remarks:                  &v.Remarks,
			Status:                   &v.Status,
		})
	}

	return resp, nil
}
