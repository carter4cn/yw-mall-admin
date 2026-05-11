package logic

import (
	"context"

	"mall-admin-api/internal/middleware"
	"mall-admin-api/internal/svc"
	"mall-admin-api/internal/types"

	"mall-shop-rpc/shopservice"
)

// ListShopApplications returns paginated shop applications for admin review.
func ListShopApplications(ctx context.Context, svcCtx *svc.ServiceContext, req *types.ListShopApplicationsReq) (*types.ListShopApplicationsResp, error) {
	resp, err := svcCtx.ShopRpc.ListShopApplications(ctx, &shopservice.ListShopApplicationsReq{
		Status:   req.Status,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	return mapApplicationsResp(resp), nil
}

func GetShopApplication(ctx context.Context, svcCtx *svc.ServiceContext, id int64) (*types.ShopApplication, error) {
	resp, err := svcCtx.ShopRpc.GetShopApplication(ctx, &shopservice.GetShopApplicationReq{Id: id})
	if err != nil {
		return nil, err
	}
	return mapApplication(resp), nil
}

func ReviewShopApplication(ctx context.Context, svcCtx *svc.ServiceContext, id int64, req *types.ReviewShopApplicationReq) (*types.OkResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	var reviewerId int64
	if c != nil {
		reviewerId = c.Uid
	}
	if _, err := svcCtx.ShopRpc.ReviewShopApplication(ctx, &shopservice.ReviewShopApplicationReq{
		ApplicationId: id,
		ReviewerId:    reviewerId,
		Action:        req.Action,
		Remark:        req.Remark,
	}); err != nil {
		return nil, err
	}
	return &types.OkResp{Ok: true}, nil
}

func ListShopsAdmin(ctx context.Context, svcCtx *svc.ServiceContext, req *types.ListShopsReq) (*types.ListShopsResp, error) {
	resp, err := svcCtx.ShopRpc.ListShops(ctx, &shopservice.ListShopsReq{Page: req.Page, PageSize: req.PageSize})
	if err != nil {
		return nil, err
	}
	out := make([]*types.ShopBrief, 0, len(resp.Shops))
	for _, s := range resp.Shops {
		out = append(out, &types.ShopBrief{
			Id:           s.Id,
			Name:         s.Name,
			Logo:         s.Logo,
			Status:       s.Status,
			CreateTime:   s.CreateTime,
			Rating:       s.Rating,
			ProductCount: s.ProductCount,
		})
	}
	return &types.ListShopsResp{Total: resp.Total, Shops: out}, nil
}

func UpdateShopStatus(ctx context.Context, svcCtx *svc.ServiceContext, shopId int64, req *types.UpdateShopStatusReq) (*types.OkResp, error) {
	if _, err := svcCtx.ShopRpc.UpdateShopStatus(ctx, &shopservice.UpdateShopStatusReq{
		ShopId: shopId,
		Status: req.Status,
		Reason: req.Reason,
	}); err != nil {
		return nil, err
	}
	return &types.OkResp{Ok: true}, nil
}

func AdjustCreditScore(ctx context.Context, svcCtx *svc.ServiceContext, shopId int64, req *types.AdjustCreditScoreReq) (*types.OkResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	var operatorId int64
	if c != nil {
		operatorId = c.Uid
	}
	if _, err := svcCtx.ShopRpc.AdjustCreditScore(ctx, &shopservice.AdjustCreditScoreReq{
		ShopId:     shopId,
		Delta:      req.Delta,
		Reason:     req.Reason,
		OperatorId: operatorId,
	}); err != nil {
		return nil, err
	}
	return &types.OkResp{Ok: true}, nil
}

func mapApplication(a *shopservice.ShopApplication) *types.ShopApplication {
	return &types.ShopApplication{
		Id:              a.Id,
		UserId:          a.UserId,
		ShopName:        a.ShopName,
		Logo:            a.Logo,
		Description:     a.Description,
		ContactPhone:    a.ContactPhone,
		BusinessLicense: a.BusinessLicense,
		LegalPerson:     a.LegalPerson,
		IdCardFront:     a.IdCardFront,
		IdCardBack:      a.IdCardBack,
		Category:        a.Category,
		Status:          a.Status,
		ReviewRemark:    a.ReviewRemark,
		ReviewerId:      a.ReviewerId,
		ShopId:          a.ShopId,
		CreateTime:      a.CreateTime,
		UpdateTime:      a.UpdateTime,
	}
}

func mapApplicationsResp(resp *shopservice.ListShopApplicationsResp) *types.ListShopApplicationsResp {
	out := make([]*types.ShopApplication, 0, len(resp.Applications))
	for _, a := range resp.Applications {
		out = append(out, mapApplication(a))
	}
	return &types.ListShopApplicationsResp{Total: resp.Total, Applications: out}
}
