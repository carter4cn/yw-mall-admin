package logic

import (
	"context"
	"errors"

	"mall-admin-api/internal/middleware"
	"mall-admin-api/internal/svc"
	"mall-admin-api/internal/types"

	"mall-order-rpc/orderclient"
)

func MerchantListOrders(ctx context.Context, svcCtx *svc.ServiceContext, req *types.MerchantListOrdersReq) (*types.MerchantListOrdersResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	if c == nil || c.ShopId <= 0 {
		return nil, errors.New("merchant shop unknown")
	}
	resp, err := svcCtx.OrderRpc.ListShopOrders(ctx, &orderclient.ListShopOrdersReq{
		ShopId:   c.ShopId,
		Status:   req.Status,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	out := make([]*types.OrderBrief, 0, len(resp.Orders))
	for _, o := range resp.Orders {
		items := make([]*types.OrderItem, 0, len(o.Items))
		for _, it := range o.Items {
			items = append(items, &types.OrderItem{
				ProductId:   it.ProductId,
				ProductName: it.ProductName,
				Price:       it.Price,
				Quantity:    it.Quantity,
			})
		}
		out = append(out, &types.OrderBrief{
			Id:          o.Id,
			OrderNo:     o.OrderNo,
			UserId:      o.UserId,
			TotalAmount: o.TotalAmount,
			Status:      o.Status,
			CreateTime:  o.CreateTime,
			Items:       items,
		})
	}
	return &types.MerchantListOrdersResp{Total: resp.Total, Orders: out}, nil
}

func MerchantGetOrder(ctx context.Context, svcCtx *svc.ServiceContext, id int64) (*types.OrderBrief, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	if c == nil || c.ShopId <= 0 {
		return nil, errors.New("merchant shop unknown")
	}
	o, err := svcCtx.OrderRpc.GetShopOrder(ctx, &orderclient.GetShopOrderReq{Id: id, ShopId: c.ShopId})
	if err != nil {
		return nil, err
	}
	items := make([]*types.OrderItem, 0, len(o.Items))
	for _, it := range o.Items {
		items = append(items, &types.OrderItem{
			ProductId:   it.ProductId,
			ProductName: it.ProductName,
			Price:       it.Price,
			Quantity:    it.Quantity,
		})
	}
	return &types.OrderBrief{
		Id:          o.Id,
		OrderNo:     o.OrderNo,
		UserId:      o.UserId,
		TotalAmount: o.TotalAmount,
		Status:      o.Status,
		CreateTime:  o.CreateTime,
		Items:       items,
	}, nil
}

func ShipOrder(ctx context.Context, svcCtx *svc.ServiceContext, id int64, req *types.ShipOrderReq) (*types.OkResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	if c == nil || c.ShopId <= 0 {
		return nil, errors.New("merchant shop unknown")
	}
	if _, err := svcCtx.OrderRpc.ShipOrder(ctx, &orderclient.ShipOrderReq{
		Id:         id,
		ShopId:     c.ShopId,
		Carrier:    req.Carrier,
		TrackingNo: req.TrackingNo,
	}); err != nil {
		return nil, err
	}
	return &types.OkResp{Ok: true}, nil
}

func MerchantRejectRefund(ctx context.Context, svcCtx *svc.ServiceContext, id int64, req *types.RejectRefundReq) (*types.OkResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	if c == nil || c.ShopId <= 0 {
		return nil, errors.New("merchant shop unknown")
	}
	if _, err := svcCtx.OrderRpc.MerchantRejectRefund(ctx, &orderclient.MerchantRejectRefundReq{
		Id:     id,
		ShopId: c.ShopId,
		Reason: req.Reason,
	}); err != nil {
		return nil, err
	}
	return &types.OkResp{Ok: true}, nil
}
