package logic

import (
	"context"

	"mall-admin-api/internal/middleware"
	"mall-admin-api/internal/svc"
	"mall-admin-api/internal/types"

	"mall-product-rpc/productclient"
)

func AdminListReviewProducts(ctx context.Context, svcCtx *svc.ServiceContext, req *types.AdminListReviewProductsReq) (*types.ListProductsResp, error) {
	resp, err := svcCtx.ProductRpc.AdminListReviewProducts(ctx, &productclient.AdminListReviewProductsReq{
		ReviewStatus: req.ReviewStatus,
		Page:         req.Page,
		PageSize:     req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	return mapProducts(resp.Products, resp.Total), nil
}

func AdminReviewProduct(ctx context.Context, svcCtx *svc.ServiceContext, id int64, req *types.AdminReviewProductReq) (*types.OkResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	var reviewerId int64
	if c != nil {
		reviewerId = c.Uid
	}
	if _, err := svcCtx.ProductRpc.AdminReviewProduct(ctx, &productclient.AdminReviewProductReq{
		Id:         id,
		Action:     req.Action,
		Remark:     req.Remark,
		ReviewerId: reviewerId,
	}); err != nil {
		return nil, err
	}
	return &types.OkResp{Ok: true}, nil
}

func mapProducts(items []*productclient.GetProductResp, total int64) *types.ListProductsResp {
	out := make([]*types.ProductBrief, 0, len(items))
	for _, p := range items {
		out = append(out, &types.ProductBrief{
			Id:          p.Id,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Stock:       p.Stock,
			Images:      p.Images,
			ShopId:      p.ShopId,
			Status:      p.Status,
			CategoryId:  p.CategoryId,
			CreateTime:  p.CreateTime,
		})
	}
	return &types.ListProductsResp{Total: total, Products: out}
}
