package emstockwatch

import (
	"context"

	"github.com/suyuan32/simple-admin-efinance-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-efinance-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-efinance-rpc/types/efinance"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmStockWatchByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEmStockWatchByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmStockWatchByIdLogic {
	return &GetEmStockWatchByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetEmStockWatchByIdLogic) GetEmStockWatchById(in *efinance.IDInt32Req) (*efinance.EmStockWatchInfo, error) {
	result, err := l.svcCtx.DB.EmStockWatch.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &efinance.EmStockWatchInfo{
		Id:                   &result.ID,
		StockWatchCode:       &result.StockWatchCode,
		StockWatchName:       &result.StockWatchName,
		StockWatchUpPrice:    &result.StockWatchUpPrice,
		StockWatchFallPrice:  &result.StockWatchFallPrice,
		StockWatchUpChange:   &result.StockWatchUpChange,
		StockWatchFallChange: &result.StockWatchFallChange,
	}, nil
}
