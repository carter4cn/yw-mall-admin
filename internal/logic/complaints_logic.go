package logic

import (
	"context"
	"errors"

	"mall-admin-api/internal/middleware"
	"mall-admin-api/internal/svc"
	"mall-admin-api/internal/types"

	"mall-risk-rpc/riskclient"
)

func mapComplaint(t *riskclient.ComplaintTicket) *types.ComplaintTicketInfo {
	return &types.ComplaintTicketInfo{
		Id:              t.Id,
		ComplainantType: t.ComplainantType,
		ComplainantId:   t.ComplainantId,
		DefendantType:   t.DefendantType,
		DefendantId:     t.DefendantId,
		OrderId:         t.OrderId,
		Category:        t.Category,
		Content:         t.Content,
		EvidenceUrls:    t.EvidenceUrls,
		Status:          t.Status,
		AdminId:         t.AdminId,
		AdminRemark:     t.AdminRemark,
		CreateTime:      t.CreateTime,
		UpdateTime:      t.UpdateTime,
	}
}

// CreateComplaint files a complaint on behalf of the calling merchant
// (complainant_type=shop, complainant_id=shop_id from JWT).
func CreateComplaint(ctx context.Context, svcCtx *svc.ServiceContext, req *types.CreateComplaintReq) (*types.CreateComplaintResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	if c == nil || c.ShopId <= 0 {
		return nil, errors.New("merchant shop unknown")
	}
	resp, err := svcCtx.RiskRpc.CreateComplaint(ctx, &riskclient.CreateComplaintReq{
		ComplainantType: "shop",
		ComplainantId:   c.ShopId,
		DefendantType:   req.DefendantType,
		DefendantId:     req.DefendantId,
		OrderId:         req.OrderId,
		Category:        req.Category,
		Content:         req.Content,
		EvidenceUrls:    req.EvidenceUrls,
	})
	if err != nil {
		return nil, err
	}
	return &types.CreateComplaintResp{Id: resp.Id}, nil
}

func ListComplaints(ctx context.Context, svcCtx *svc.ServiceContext, req *types.ListComplaintsReq) (*types.ListComplaintsResp, error) {
	resp, err := svcCtx.RiskRpc.ListComplaints(ctx, &riskclient.ListComplaintsReq{
		Status:        req.Status,
		DefendantType: req.DefendantType,
		DefendantId:   req.DefendantId,
		Page:          req.Page,
		PageSize:      req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	out := make([]*types.ComplaintTicketInfo, 0, len(resp.Tickets))
	for _, t := range resp.Tickets {
		out = append(out, mapComplaint(t))
	}
	return &types.ListComplaintsResp{Total: resp.Total, Tickets: out}, nil
}

func GetComplaint(ctx context.Context, svcCtx *svc.ServiceContext, id int64) (*types.ComplaintTicketInfo, error) {
	t, err := svcCtx.RiskRpc.GetComplaint(ctx, &riskclient.GetComplaintReq{Id: id})
	if err != nil {
		return nil, err
	}
	return mapComplaint(t), nil
}

func HandleComplaint(ctx context.Context, svcCtx *svc.ServiceContext, id int64, req *types.HandleComplaintReq) (*types.OkResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	var adminId int64
	if c != nil {
		adminId = c.Uid
	}
	if _, err := svcCtx.RiskRpc.HandleComplaint(ctx, &riskclient.HandleComplaintReq{
		Id:      id,
		AdminId: adminId,
		Action:  req.Action,
		Remark:  req.Remark,
	}); err != nil {
		return nil, err
	}
	return &types.OkResp{Ok: true}, nil
}

func SetShopRestriction(ctx context.Context, svcCtx *svc.ServiceContext, shopId int64, req *types.SetShopRestrictionReq) (*types.OkResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	var operatorId int64
	if c != nil {
		operatorId = c.Uid
	}
	if _, err := svcCtx.RiskRpc.SetShopRestriction(ctx, &riskclient.SetShopRestrictionReq{
		ShopId:      shopId,
		Restriction: req.Restriction,
		Reason:      req.Reason,
		OperatorId:  operatorId,
		ExpireTime:  req.ExpireTime,
	}); err != nil {
		return nil, err
	}
	return &types.OkResp{Ok: true}, nil
}

func ListShopRestrictions(ctx context.Context, svcCtx *svc.ServiceContext, shopId int64) (*types.ListShopRestrictionsResp, error) {
	resp, err := svcCtx.RiskRpc.ListShopRestrictions(ctx, &riskclient.ListShopRestrictionsReq{ShopId: shopId})
	if err != nil {
		return nil, err
	}
	out := make([]*types.ShopRestrictionInfo, 0, len(resp.Restrictions))
	for _, r := range resp.Restrictions {
		out = append(out, &types.ShopRestrictionInfo{
			Id:          r.Id,
			ShopId:      r.ShopId,
			Restriction: r.Restriction,
			Reason:      r.Reason,
			OperatorId:  r.OperatorId,
			ExpireTime:  r.ExpireTime,
			CreateTime:  r.CreateTime,
		})
	}
	return &types.ListShopRestrictionsResp{Restrictions: out}, nil
}

func RemoveShopRestriction(ctx context.Context, svcCtx *svc.ServiceContext, restrictionId int64) (*types.OkResp, error) {
	if _, err := svcCtx.RiskRpc.RemoveShopRestriction(ctx, &riskclient.RemoveShopRestrictionReq{Id: restrictionId}); err != nil {
		return nil, err
	}
	return &types.OkResp{Ok: true}, nil
}
