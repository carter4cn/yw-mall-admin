package logic

import (
	"context"
	"errors"

	"mall-admin-api/internal/middleware"
	"mall-admin-api/internal/svc"
	"mall-admin-api/internal/types"

	"mall-shop-rpc/shopservice"
)

func ApplyShop(ctx context.Context, svcCtx *svc.ServiceContext, req *types.ApplyShopReq) (*types.ApplyShopResp, error) {
	resp, err := svcCtx.ShopRpc.ApplyShop(ctx, &shopservice.ApplyShopReq{
		UserId:          req.UserId,
		ShopName:        req.ShopName,
		Logo:            req.Logo,
		Description:     req.Description,
		ContactPhone:    req.ContactPhone,
		BusinessLicense: req.BusinessLicense,
		LegalPerson:     req.LegalPerson,
		IdCardFront:     req.IdCardFront,
		IdCardBack:      req.IdCardBack,
		Category:        req.Category,
	})
	if err != nil {
		return nil, err
	}
	return &types.ApplyShopResp{ApplicationId: resp.ApplicationId}, nil
}

func GetMyApplication(ctx context.Context, svcCtx *svc.ServiceContext, id int64) (*types.ShopApplication, error) {
	return GetShopApplication(ctx, svcCtx, id)
}

func GetMyShop(ctx context.Context, svcCtx *svc.ServiceContext) (*types.ShopDetail, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	if c == nil || c.Uid <= 0 {
		return nil, errors.New("unauthorized")
	}
	resp, err := svcCtx.ShopRpc.GetShopByOwnerId(ctx, &shopservice.GetShopByOwnerIdReq{OwnerUserId: c.Uid})
	if err != nil {
		return nil, err
	}
	return &types.ShopDetail{
		Id:              resp.Id,
		Name:            resp.Name,
		Logo:            resp.Logo,
		Banner:          resp.Banner,
		Description:     resp.Description,
		Rating:          resp.Rating,
		ProductCount:    resp.ProductCount,
		FollowCount:     resp.FollowCount,
		Status:          resp.Status,
		CreateTime:      resp.CreateTime,
		OwnerUserId:     resp.OwnerUserId,
		CreditScore:     resp.CreditScore,
		Level:           resp.Level,
		ContactPhone:    resp.ContactPhone,
		BusinessLicense: resp.BusinessLicense,
	}, nil
}

func UpdateMyShop(ctx context.Context, svcCtx *svc.ServiceContext, req *types.UpdateMyShopReq) (*types.OkResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	if c == nil || c.ShopId <= 0 {
		return nil, errors.New("merchant shop unknown")
	}
	if _, err := svcCtx.ShopRpc.UpdateShop(ctx, &shopservice.UpdateShopReq{
		Id:          c.ShopId,
		Name:        req.Name,
		Logo:        req.Logo,
		Banner:      req.Banner,
		Description: req.Description,
	}); err != nil {
		return nil, err
	}
	return &types.OkResp{Ok: true}, nil
}
