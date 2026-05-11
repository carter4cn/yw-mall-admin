package logic

import (
	"context"
	"errors"

	"mall-admin-api/internal/middleware"
	"mall-admin-api/internal/svc"
	"mall-admin-api/internal/types"

	"mall-shop-rpc/shopservice"
)

// SubmitShopLifecycleRequest lets the current merchant request a lifecycle
// state change (deactivate / pause / resume) on their own shop.
func SubmitShopLifecycleRequest(ctx context.Context, svcCtx *svc.ServiceContext, req *types.SubmitShopLifecycleRequestReq) (*types.SubmitShopLifecycleRequestResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	if c == nil || c.ShopId <= 0 {
		return nil, errors.New("merchant shop unknown")
	}
	resp, err := svcCtx.ShopRpc.SubmitShopLifecycleRequest(ctx, &shopservice.SubmitShopLifecycleRequestReq{
		ShopId: c.ShopId,
		Action: req.Action,
		Reason: req.Reason,
	})
	if err != nil {
		return nil, err
	}
	return &types.SubmitShopLifecycleRequestResp{RequestId: resp.RequestId}, nil
}

// ListShopLifecycleRequests returns paginated lifecycle requests for admin.
func ListShopLifecycleRequests(ctx context.Context, svcCtx *svc.ServiceContext, req *types.ListShopLifecycleRequestsReq) (*types.ListShopLifecycleRequestsResp, error) {
	resp, err := svcCtx.ShopRpc.ListShopLifecycleRequests(ctx, &shopservice.ListShopLifecycleRequestsReq{
		Status:   req.Status,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	out := make([]*types.ShopLifecycleRequestInfo, 0, len(resp.Requests))
	for _, r := range resp.Requests {
		out = append(out, &types.ShopLifecycleRequestInfo{
			Id:          r.Id,
			ShopId:      r.ShopId,
			Action:      r.Action,
			Reason:      r.Reason,
			Status:      r.Status,
			AdminId:     r.AdminId,
			AdminRemark: r.AdminRemark,
			CreateTime:  r.CreateTime,
		})
	}
	return &types.ListShopLifecycleRequestsResp{Total: resp.Total, Requests: out}, nil
}

// ReviewShopLifecycleRequest approves or rejects a lifecycle change.
func ReviewShopLifecycleRequest(ctx context.Context, svcCtx *svc.ServiceContext, id int64, req *types.ReviewShopLifecycleRequestReq) (*types.OkResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	var adminId int64
	if c != nil {
		adminId = c.Uid
	}
	if _, err := svcCtx.ShopRpc.ReviewShopLifecycleRequest(ctx, &shopservice.ReviewShopLifecycleRequestReq{
		RequestId: id,
		AdminId:   adminId,
		Action:    req.Action,
		Remark:    req.Remark,
	}); err != nil {
		return nil, err
	}
	return &types.OkResp{Ok: true}, nil
}
