package emstockbasicinfo

import (
	"context"

    "github.com/suyuan32/simple-admin-efinance-rpc/ent/emstockbasicinfo"
    "github.com/suyuan32/simple-admin-efinance-rpc/internal/svc"
    "github.com/suyuan32/simple-admin-efinance-rpc/internal/utils/dberrorhandler"
    "github.com/suyuan32/simple-admin-efinance-rpc/types/efinance"

    "github.com/suyuan32/simple-admin-common/msg/errormsg"
    "github.com/zeromicro/go-zero/core/logx"
)

type DeleteEmStockBasicInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteEmStockBasicInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteEmStockBasicInfoLogic {
	return &DeleteEmStockBasicInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteEmStockBasicInfoLogic) DeleteEmStockBasicInfo(in *efinance.IDsInt32Req) (*efinance.BaseResp, error) {
	_, err := l.svcCtx.DB.EmStockBasicInfo.Delete().Where(emstockbasicinfo.IDIn(in.Ids...)).Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &efinance.BaseResp{Msg: errormsg.DeleteSuccess }, nil
}
