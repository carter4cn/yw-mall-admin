package logic

import (
	"context"
	"errors"

	"mall-admin-api/internal/middleware"
	"mall-admin-api/internal/svc"
	"mall-admin-api/internal/types"

	"mall-review-rpc/reviewclient"
)

// MerchantListReviews returns reviews scoped to the calling merchant's shop.
func MerchantListReviews(ctx context.Context, svcCtx *svc.ServiceContext, req *types.MerchantListReviewsReq) (*types.ListReviewsResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	if c == nil || c.ShopId <= 0 {
		return nil, errors.New("merchant shop unknown")
	}
	resp, err := svcCtx.ReviewRpc.ListShopReviews(ctx, &reviewclient.ListShopReviewsReq{
		ShopId:   c.ShopId,
		Score:    req.Score,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	out := make([]*types.ReviewBrief, 0, len(resp.Reviews))
	for _, r := range resp.Reviews {
		out = append(out, &types.ReviewBrief{
			Id:                r.Id,
			OrderItemId:       r.OrderItemId,
			UserId:            r.UserId,
			ProductId:         r.ProductId,
			ScoreOverall:      r.ScoreOverall,
			Content:           r.Content,
			MerchantReplyText: r.MerchantReplyText,
			Status:            r.Status,
			CreateTime:        r.CreateTime,
		})
	}
	return &types.ListReviewsResp{Total: resp.Total, Reviews: out}, nil
}

func RequestDeleteReview(ctx context.Context, svcCtx *svc.ServiceContext, reviewId int64, req *types.RequestDeleteReviewReq) (*types.OkResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	if c == nil || c.ShopId <= 0 {
		return nil, errors.New("merchant shop unknown")
	}
	if _, err := svcCtx.ReviewRpc.RequestDeleteReview(ctx, &reviewclient.RequestDeleteReviewReq{
		ReviewId: reviewId,
		ShopId:   c.ShopId,
		Reason:   req.Reason,
	}); err != nil {
		return nil, err
	}
	return &types.OkResp{Ok: true}, nil
}

func ListDeleteRequests(ctx context.Context, svcCtx *svc.ServiceContext, req *types.ListDeleteRequestsReq) (*types.ListDeleteRequestsResp, error) {
	resp, err := svcCtx.ReviewRpc.ListDeleteRequests(ctx, &reviewclient.ListDeleteRequestsReq{
		Status:   req.Status,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	out := make([]*types.ReviewDeleteRequestInfo, 0, len(resp.Requests))
	for _, r := range resp.Requests {
		out = append(out, &types.ReviewDeleteRequestInfo{
			Id:          r.Id,
			ReviewId:    r.ReviewId,
			ShopId:      r.ShopId,
			Reason:      r.Reason,
			Status:      r.Status,
			AdminRemark: r.AdminRemark,
			AdminId:     r.AdminId,
			CreateTime:  r.CreateTime,
		})
	}
	return &types.ListDeleteRequestsResp{Total: resp.Total, Requests: out}, nil
}

func AdminHandleDeleteRequest(ctx context.Context, svcCtx *svc.ServiceContext, requestId int64, req *types.AdminHandleDeleteRequestReq) (*types.OkResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	var adminId int64
	if c != nil {
		adminId = c.Uid
	}
	if _, err := svcCtx.ReviewRpc.AdminHandleDeleteRequest(ctx, &reviewclient.AdminHandleDeleteRequestReq{
		RequestId: requestId,
		AdminId:   adminId,
		Action:    req.Action,
		Remark:    req.Remark,
	}); err != nil {
		return nil, err
	}
	return &types.OkResp{Ok: true}, nil
}
