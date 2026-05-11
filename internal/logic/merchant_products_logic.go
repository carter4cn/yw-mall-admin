package logic

import (
	"context"
	"errors"

	"mall-admin-api/internal/middleware"
	"mall-admin-api/internal/svc"
	"mall-admin-api/internal/types"

	"mall-product-rpc/productclient"
)

func MerchantListProducts(ctx context.Context, svcCtx *svc.ServiceContext, req *types.MerchantListProductsReq) (*types.ListProductsResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	if c == nil || c.ShopId <= 0 {
		return nil, errors.New("merchant shop unknown")
	}
	resp, err := svcCtx.ProductRpc.MerchantListProducts(ctx, &productclient.MerchantListProductsReq{
		ShopId:       c.ShopId,
		Status:       req.Status,
		ReviewStatus: req.ReviewStatus,
		Page:         req.Page,
		PageSize:     req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	return mapProducts(resp.Products, resp.Total), nil
}

func MerchantCreateProduct(ctx context.Context, svcCtx *svc.ServiceContext, req *types.MerchantCreateProductReq) (*types.MerchantCreateProductResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	if c == nil || c.ShopId <= 0 {
		return nil, errors.New("merchant shop unknown")
	}
	resp, err := svcCtx.ProductRpc.CreateProduct(ctx, &productclient.CreateProductReq{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		CategoryId:  req.CategoryId,
		Images:      req.Images,
		ShopId:      c.ShopId,
	})
	if err != nil {
		return nil, err
	}
	// optional rich fields via UpdateProduct (no-op if values are zero); the merchant
	// should call PUT /products/:id afterwards if they want to set detail/brand/weight.
	if req.Detail != "" || req.Brand != "" || req.Weight > 0 {
		_, _ = svcCtx.ProductRpc.UpdateProduct(ctx, &productclient.UpdateProductReq{
			Id:          resp.Id,
			ShopId:      c.ShopId,
			Name:        req.Name,
			Description: req.Description,
			Price:       req.Price,
			CategoryId:  req.CategoryId,
			Images:      req.Images,
			Detail:      req.Detail,
			Brand:       req.Brand,
			Weight:      req.Weight,
		})
	}
	return &types.MerchantCreateProductResp{Id: resp.Id}, nil
}

func MerchantGetProduct(ctx context.Context, svcCtx *svc.ServiceContext, id int64) (*types.ProductDetail, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	if c == nil {
		return nil, errors.New("unauthorized")
	}
	resp, err := svcCtx.ProductRpc.GetProductDetail(ctx, &productclient.GetProductReq{Id: id})
	if err != nil {
		return nil, err
	}
	if c.ShopId > 0 && resp.ShopId != c.ShopId {
		return nil, errors.New("product not owned by shop")
	}
	return &types.ProductDetail{
		Id:           resp.Id,
		Name:         resp.Name,
		Description:  resp.Description,
		Price:        resp.Price,
		Stock:        resp.Stock,
		CategoryId:   resp.CategoryId,
		Images:       resp.Images,
		Status:       resp.Status,
		ShopId:       resp.ShopId,
		ReviewStatus: resp.ReviewStatus,
		ReviewRemark: resp.ReviewRemark,
		Detail:       resp.Detail,
		Brand:        resp.Brand,
		Weight:       resp.Weight,
		CreateTime:   resp.CreateTime,
	}, nil
}

func MerchantUpdateProduct(ctx context.Context, svcCtx *svc.ServiceContext, id int64, req *types.MerchantUpdateProductReq) (*types.OkResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	if c == nil || c.ShopId <= 0 {
		return nil, errors.New("merchant shop unknown")
	}
	if _, err := svcCtx.ProductRpc.UpdateProduct(ctx, &productclient.UpdateProductReq{
		Id:          id,
		ShopId:      c.ShopId,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		CategoryId:  req.CategoryId,
		Images:      req.Images,
		Detail:      req.Detail,
		Brand:       req.Brand,
		Weight:      req.Weight,
	}); err != nil {
		return nil, err
	}
	return &types.OkResp{Ok: true}, nil
}

func MerchantSetProductStatus(ctx context.Context, svcCtx *svc.ServiceContext, id int64, req *types.SetProductStatusReq) (*types.OkResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	if c == nil || c.ShopId <= 0 {
		return nil, errors.New("merchant shop unknown")
	}
	if _, err := svcCtx.ProductRpc.SetProductStatus(ctx, &productclient.SetProductStatusReq{
		Id:     id,
		ShopId: c.ShopId,
		Status: req.Status,
	}); err != nil {
		return nil, err
	}
	return &types.OkResp{Ok: true}, nil
}

func MerchantSetProductStock(ctx context.Context, svcCtx *svc.ServiceContext, id int64, req *types.SetProductStockReq) (*types.OkResp, error) {
	if _, err := svcCtx.ProductRpc.SetProductStock(ctx, &productclient.SetProductStockReq{
		Id:    id,
		Stock: req.Stock,
	}); err != nil {
		return nil, err
	}
	return &types.OkResp{Ok: true}, nil
}
