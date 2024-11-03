package emstockwatch

import (
	"context"

	"github.com/suyuan32/simple-admin-efinance-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-efinance-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-efinance-rpc/types/efinance"

    "github.com/suyuan32/simple-admin-common/msg/errormsg"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateEmStockWatchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateEmStockWatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateEmStockWatchLogic {
	return &UpdateEmStockWatchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateEmStockWatchLogic) UpdateEmStockWatch(in *efinance.EmStockWatchInfo) (*efinance.BaseResp, error) {
	err:= l.svcCtx.DB.EmStockWatch.UpdateOneID(*in.Id).
			SetNotNilStockWatchCode(in.StockWatchCode).
			SetNotNilStockWatchName(in.StockWatchName).
			SetNotNilStockWatchUpPrice(in.StockWatchUpPrice).
			SetNotNilStockWatchFallPrice(in.StockWatchFallPrice).
			SetNotNilStockWatchUpChange(in.StockWatchUpChange).
			SetNotNilStockWatchFallChange(in.StockWatchFallChange).
			Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &efinance.BaseResp{Msg: errormsg.UpdateSuccess }, nil
}
