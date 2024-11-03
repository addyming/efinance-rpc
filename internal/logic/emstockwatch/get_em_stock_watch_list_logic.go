package emstockwatch

import (
	"context"

	"github.com/suyuan32/simple-admin-efinance-rpc/ent/emstockwatch"
	"github.com/suyuan32/simple-admin-efinance-rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-efinance-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-efinance-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-efinance-rpc/types/efinance"


    "github.com/zeromicro/go-zero/core/logx"
)

type GetEmStockWatchListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEmStockWatchListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmStockWatchListLogic {
	return &GetEmStockWatchListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetEmStockWatchListLogic) GetEmStockWatchList(in *efinance.EmStockWatchListReq) (*efinance.EmStockWatchListResp, error) {
	var predicates []predicate.EmStockWatch
	if in.StockWatchCode != nil {
		predicates = append(predicates, emstockwatch.StockWatchCodeContains(*in.StockWatchCode))
	}
	if in.StockWatchName != nil {
		predicates = append(predicates, emstockwatch.StockWatchNameContains(*in.StockWatchName))
	}
	if in.StockWatchUpPrice != nil {
		predicates = append(predicates, emstockwatch.StockWatchUpPriceContains(*in.StockWatchUpPrice))
	}
	result, err := l.svcCtx.DB.EmStockWatch.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &efinance.EmStockWatchListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &efinance.EmStockWatchInfo{
			Id:          &v.ID,
			StockWatchCode:	&v.StockWatchCode,
			StockWatchName:	&v.StockWatchName,
			StockWatchUpPrice:	&v.StockWatchUpPrice,
			StockWatchFallPrice:	&v.StockWatchFallPrice,
			StockWatchUpChange:	&v.StockWatchUpChange,
			StockWatchFallChange:	&v.StockWatchFallChange,
		})
	}

	return resp, nil
}
