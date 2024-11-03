package emstockwatch

import (
	"context"

    "github.com/suyuan32/simple-admin-efinance-rpc/ent/emstockwatch"
    "github.com/suyuan32/simple-admin-efinance-rpc/internal/svc"
    "github.com/suyuan32/simple-admin-efinance-rpc/internal/utils/dberrorhandler"
    "github.com/suyuan32/simple-admin-efinance-rpc/types/efinance"

    "github.com/suyuan32/simple-admin-common/msg/errormsg"
    "github.com/zeromicro/go-zero/core/logx"
)

type DeleteEmStockWatchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteEmStockWatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteEmStockWatchLogic {
	return &DeleteEmStockWatchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteEmStockWatchLogic) DeleteEmStockWatch(in *efinance.IDsInt32Req) (*efinance.BaseResp, error) {
	_, err := l.svcCtx.DB.EmStockWatch.Delete().Where(emstockwatch.IDIn(in.Ids...)).Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &efinance.BaseResp{Msg: errormsg.DeleteSuccess }, nil
}
