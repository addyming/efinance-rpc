package emstockbasicinfo

import (
	"context"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-efinance-rpc/ent/emstockbasicinfo"
	"github.com/suyuan32/simple-admin-efinance-rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-efinance-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-efinance-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-efinance-rpc/types/efinance"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmStockBasicInfoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEmStockBasicInfoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmStockBasicInfoListLogic {
	return &GetEmStockBasicInfoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetEmStockBasicInfoListLogic) GetEmStockBasicInfoList(in *efinance.EmStockBasicInfoListReq) (*efinance.EmStockBasicInfoListResp, error) {
	var predicates []predicate.EmStockBasicInfo
	if in.F1 != nil {
		predicates = append(predicates, emstockbasicinfo.F1Contains(*in.F1))
	}
	if in.F2 != nil {
		predicates = append(predicates, emstockbasicinfo.F2Contains(*in.F2))
	}
	if in.F3 != nil {
		predicates = append(predicates, emstockbasicinfo.F3Contains(*in.F3))
	}
	result, err := l.svcCtx.DB.EmStockBasicInfo.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &efinance.EmStockBasicInfoListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &efinance.EmStockBasicInfoInfo{
			Id:         &v.ID,
			F1:         &v.F1,
			F2:         &v.F2,
			F3:         &v.F3,
			F4:         &v.F4,
			F5:         &v.F5,
			F6:         &v.F6,
			F7:         &v.F7,
			F8:         &v.F8,
			F9:         &v.F9,
			F10:        &v.F10,
			F11:        &v.F11,
			F12:        &v.F12,
			F13:        &v.F13,
			F14:        &v.F14,
			F15:        &v.F15,
			F16:        &v.F16,
			F17:        &v.F17,
			F18:        &v.F18,
			F19:        &v.F19,
			F20:        &v.F20,
			F21:        &v.F21,
			F22:        &v.F22,
			F23:        &v.F23,
			F24:        &v.F24,
			F25:        &v.F25,
			F26:        &v.F26,
			F27:        &v.F27,
			F28:        &v.F28,
			F29:        &v.F29,
			F30:        &v.F30,
			CreateTime: pointy.GetUnixMilliPointer(v.CreateTime.UnixMilli()),
			Remarks:    &v.Remarks,
			Status:     &v.Status,
			St1:        &v.St1,
			St2:        &v.St2,
			St3:        &v.St3,
			St4:        &v.St4,
			St5:        &v.St5,
		})
	}

	return resp, nil
}
