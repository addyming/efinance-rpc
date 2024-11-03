package emstockbasicinfo

import (
	"context"

	"github.com/suyuan32/simple-admin-efinance-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-efinance-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-efinance-rpc/types/efinance"

    "github.com/suyuan32/simple-admin-common/msg/errormsg"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateEmStockBasicInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateEmStockBasicInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateEmStockBasicInfoLogic {
	return &UpdateEmStockBasicInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateEmStockBasicInfoLogic) UpdateEmStockBasicInfo(in *efinance.EmStockBasicInfoInfo) (*efinance.BaseResp, error) {
	err:= l.svcCtx.DB.EmStockBasicInfo.UpdateOneID(*in.Id).
			SetNotNilF1(in.F1).
			SetNotNilF2(in.F2).
			SetNotNilF3(in.F3).
			SetNotNilF4(in.F4).
			SetNotNilF5(in.F5).
			SetNotNilF6(in.F6).
			SetNotNilF7(in.F7).
			SetNotNilF8(in.F8).
			SetNotNilF9(in.F9).
			SetNotNilF10(in.F10).
			SetNotNilF11(in.F11).
			SetNotNilF12(in.F12).
			SetNotNilF13(in.F13).
			SetNotNilF14(in.F14).
			SetNotNilF15(in.F15).
			SetNotNilF16(in.F16).
			SetNotNilF17(in.F17).
			SetNotNilF18(in.F18).
			SetNotNilF19(in.F19).
			SetNotNilF20(in.F20).
			SetNotNilF21(in.F21).
			SetNotNilF22(in.F22).
			SetNotNilF23(in.F23).
			SetNotNilF24(in.F24).
			SetNotNilF25(in.F25).
			SetNotNilF26(in.F26).
			SetNotNilF27(in.F27).
			SetNotNilF28(in.F28).
			SetNotNilF29(in.F29).
			SetNotNilF30(in.F30).
			SetNotNilCreateTime(pointy.GetTimeMilliPointer(in.CreateTime)).
			SetNotNilRemarks(in.Remarks).
			SetNotNilStatus(in.Status).
			SetNotNilSt1(in.St1).
			SetNotNilSt2(in.St2).
			SetNotNilSt3(in.St3).
			SetNotNilSt4(in.St4).
			SetNotNilSt5(in.St5).
			Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &efinance.BaseResp{Msg: errormsg.UpdateSuccess }, nil
}
