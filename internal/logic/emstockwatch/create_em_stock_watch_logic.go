package emstockwatch

import (
	"context"

	"github.com/suyuan32/simple-admin-efinance-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-efinance-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-efinance-rpc/types/efinance"

    "github.com/suyuan32/simple-admin-common/msg/errormsg"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateEmStockWatchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateEmStockWatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateEmStockWatchLogic {
	return &CreateEmStockWatchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateEmStockWatchLogic) CreateEmStockWatch(in *efinance.EmStockWatchInfo) (*efinance.BaseIDInt32Resp, error) {
    result, err := l.svcCtx.DB.EmStockWatch.Create().
			SetNotNilStockWatchCode(in.StockWatchCode).
			SetNotNilStockWatchName(in.StockWatchName).
			SetNotNilStockWatchUpPrice(in.StockWatchUpPrice).
			SetNotNilStockWatchFallPrice(in.StockWatchFallPrice).
			SetNotNilStockWatchUpChange(in.StockWatchUpChange).
			SetNotNilStockWatchFallChange(in.StockWatchFallChange).
			Save(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &efinance.BaseIDInt32Resp{Id: result.ID, Msg: errormsg.CreateSuccess }, nil
}
