package logic

import (
	"context"
	"errors"

	"mall-admin-api/internal/middleware"
	"mall-admin-api/internal/svc"
	"mall-admin-api/internal/types"

	"mall-activity-rpc/activityclient"
)

// ===== P2 G-3: Shop coupons =====

func CreateShopCoupon(ctx context.Context, svcCtx *svc.ServiceContext, req *types.CreateShopCouponReq) (*types.CreateShopCouponResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	if c == nil || c.ShopId <= 0 {
		return nil, errors.New("merchant shop unknown")
	}
	resp, err := svcCtx.ActivityRpc.CreateShopCoupon(ctx, &activityclient.CreateShopCouponReq{
		ShopId:         c.ShopId,
		Code:           req.Code,
		Name:           req.Name,
		Type:           req.Type,
		DiscountValue:  req.DiscountValue,
		MinOrderAmount: req.MinOrderAmount,
		TotalQuantity:  req.TotalQuantity,
		PerUserLimit:   req.PerUserLimit,
		ValidFrom:      req.ValidFrom,
		ValidTo:        req.ValidTo,
	})
	if err != nil {
		return nil, err
	}
	return &types.CreateShopCouponResp{Id: resp.Id}, nil
}

func ListShopCoupons(ctx context.Context, svcCtx *svc.ServiceContext, req *types.ListShopCouponsReq) (*types.ListShopCouponsResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	if c == nil || c.ShopId <= 0 {
		return nil, errors.New("merchant shop unknown")
	}
	resp, err := svcCtx.ActivityRpc.ListShopCoupons(ctx, &activityclient.ListShopCouponsReq{
		ShopId:   c.ShopId,
		Status:   req.Status,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	out := make([]*types.ShopCouponInfo, 0, len(resp.Coupons))
	for _, r := range resp.Coupons {
		out = append(out, &types.ShopCouponInfo{
			Id:              r.Id,
			ShopId:          r.ShopId,
			Code:            r.Code,
			Name:            r.Name,
			Type:            r.Type,
			DiscountValue:   r.DiscountValue,
			MinOrderAmount:  r.MinOrderAmount,
			TotalQuantity:   r.TotalQuantity,
			ClaimedQuantity: r.ClaimedQuantity,
			PerUserLimit:    r.PerUserLimit,
			ValidFrom:       r.ValidFrom,
			ValidTo:         r.ValidTo,
			Status:          r.Status,
			CreateTime:      r.CreateTime,
		})
	}
	return &types.ListShopCouponsResp{Total: resp.Total, Coupons: out}, nil
}

func UpdateShopCouponStatus(ctx context.Context, svcCtx *svc.ServiceContext, id int64, req *types.UpdateShopCouponStatusReq) (*types.OkResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	if c == nil || c.ShopId <= 0 {
		return nil, errors.New("merchant shop unknown")
	}
	if _, err := svcCtx.ActivityRpc.UpdateShopCouponStatus(ctx, &activityclient.UpdateShopCouponStatusReq{
		Id:     id,
		ShopId: c.ShopId,
		Status: req.Status,
	}); err != nil {
		return nil, err
	}
	return &types.OkResp{Ok: true}, nil
}

// ===== P2 G-4: SKU flash discounts =====

func CreateFlashDiscount(ctx context.Context, svcCtx *svc.ServiceContext, req *types.CreateFlashDiscountReq) (*types.CreateFlashDiscountResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	if c == nil || c.ShopId <= 0 {
		return nil, errors.New("merchant shop unknown")
	}
	resp, err := svcCtx.ActivityRpc.CreateFlashDiscount(ctx, &activityclient.CreateFlashDiscountReq{
		ShopId:        c.ShopId,
		ProductId:     req.ProductId,
		SkuId:         req.SkuId,
		OriginalPrice: req.OriginalPrice,
		DiscountPrice: req.DiscountPrice,
		StartTime:     req.StartTime,
		EndTime:       req.EndTime,
	})
	if err != nil {
		return nil, err
	}
	return &types.CreateFlashDiscountResp{Id: resp.Id}, nil
}

func ListFlashDiscounts(ctx context.Context, svcCtx *svc.ServiceContext, req *types.ListFlashDiscountsReq) (*types.ListFlashDiscountsResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	if c == nil || c.ShopId <= 0 {
		return nil, errors.New("merchant shop unknown")
	}
	resp, err := svcCtx.ActivityRpc.ListFlashDiscounts(ctx, &activityclient.ListFlashDiscountsReq{
		ShopId:   c.ShopId,
		Status:   req.Status,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	out := make([]*types.FlashDiscountInfo, 0, len(resp.Discounts))
	for _, r := range resp.Discounts {
		out = append(out, &types.FlashDiscountInfo{
			Id:            r.Id,
			ShopId:        r.ShopId,
			ProductId:     r.ProductId,
			SkuId:         r.SkuId,
			OriginalPrice: r.OriginalPrice,
			DiscountPrice: r.DiscountPrice,
			StartTime:     r.StartTime,
			EndTime:       r.EndTime,
			Status:        r.Status,
			CreateTime:    r.CreateTime,
		})
	}
	return &types.ListFlashDiscountsResp{Total: resp.Total, Discounts: out}, nil
}

func CancelFlashDiscount(ctx context.Context, svcCtx *svc.ServiceContext, id int64) (*types.OkResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	if c == nil || c.ShopId <= 0 {
		return nil, errors.New("merchant shop unknown")
	}
	if _, err := svcCtx.ActivityRpc.CancelFlashDiscount(ctx, &activityclient.CancelFlashDiscountReq{
		Id:     id,
		ShopId: c.ShopId,
	}); err != nil {
		return nil, err
	}
	return &types.OkResp{Ok: true}, nil
}
