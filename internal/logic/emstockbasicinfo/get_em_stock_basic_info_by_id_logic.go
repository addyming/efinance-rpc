package emstockbasicinfo

import (
	"context"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-efinance-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-efinance-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-efinance-rpc/types/efinance"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmStockBasicInfoByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEmStockBasicInfoByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmStockBasicInfoByIdLogic {
	return &GetEmStockBasicInfoByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetEmStockBasicInfoByIdLogic) GetEmStockBasicInfoById(in *efinance.IDInt32Req) (*efinance.EmStockBasicInfoInfo, error) {
	result, err := l.svcCtx.DB.EmStockBasicInfo.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &efinance.EmStockBasicInfoInfo{
		Id:         &result.ID,
		F1:         &result.F1,
		F2:         &result.F2,
		F3:         &result.F3,
		F4:         &result.F4,
		F5:         &result.F5,
		F6:         &result.F6,
		F7:         &result.F7,
		F8:         &result.F8,
		F9:         &result.F9,
		F10:        &result.F10,
		F11:        &result.F11,
		F12:        &result.F12,
		F13:        &result.F13,
		F14:        &result.F14,
		F15:        &result.F15,
		F16:        &result.F16,
		F17:        &result.F17,
		F18:        &result.F18,
		F19:        &result.F19,
		F20:        &result.F20,
		F21:        &result.F21,
		F22:        &result.F22,
		F23:        &result.F23,
		F24:        &result.F24,
		F25:        &result.F25,
		F26:        &result.F26,
		F27:        &result.F27,
		F28:        &result.F28,
		F29:        &result.F29,
		F30:        &result.F30,
		CreateTime: pointy.GetUnixMilliPointer(result.CreateTime.UnixMilli()),
		Remarks:    &result.Remarks,
		Status:     &result.Status,
		St1:        &result.St1,
		St2:        &result.St2,
		St3:        &result.St3,
		St4:        &result.St4,
		St5:        &result.St5,
	}, nil
}
