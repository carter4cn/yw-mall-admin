package logic

import (
	"context"
	"errors"

	"mall-admin-api/internal/middleware"
	"mall-admin-api/internal/svc"
	"mall-admin-api/internal/types"

	"mall-shop-rpc/shopservice"
)

// ListShopLevels returns all 5 level templates (merchant view).
func ListShopLevels(ctx context.Context, svcCtx *svc.ServiceContext) (*types.ListShopLevelsResp, error) {
	resp, err := svcCtx.ShopRpc.ListShopLevels(ctx, &shopservice.Empty{})
	if err != nil {
		return nil, err
	}
	out := make([]*types.ShopLevelTemplateInfo, 0, len(resp.Levels))
	for _, t := range resp.Levels {
		out = append(out, mapLevelTemplate(t))
	}
	return &types.ListShopLevelsResp{Levels: out}, nil
}

// GetMyLevelStatus returns current shop's progress against next level.
func GetMyLevelStatus(ctx context.Context, svcCtx *svc.ServiceContext) (*types.MyLevelStatusResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	if c == nil || c.ShopId <= 0 {
		return nil, errors.New("merchant shop unknown")
	}
	resp, err := svcCtx.ShopRpc.GetMyLevelStatus(ctx, &shopservice.GetMyLevelStatusReq{ShopId: c.ShopId})
	if err != nil {
		return nil, err
	}
	return &types.MyLevelStatusResp{
		CurrentLevel:          resp.CurrentLevel,
		CurrentTemplate:       mapLevelTemplate(resp.CurrentTemplate),
		NextTemplate:          mapLevelTemplate(resp.NextTemplate),
		CurrentGmv:            resp.CurrentGmv,
		CurrentCreditScore:    resp.CurrentCreditScore,
		CurrentMonths:         resp.CurrentMonths,
		CurrentRating:         resp.CurrentRating,
		EligibleForNext:       resp.EligibleForNext,
		HasPendingApplication: resp.HasPendingApplication,
	}, nil
}

// SubmitLevelApplication submits a level upgrade request for the current shop.
func SubmitLevelApplication(ctx context.Context, svcCtx *svc.ServiceContext, req *types.SubmitLevelApplicationReq) (*types.SubmitLevelApplicationResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	if c == nil || c.ShopId <= 0 {
		return nil, errors.New("merchant shop unknown")
	}
	resp, err := svcCtx.ShopRpc.SubmitLevelApplication(ctx, &shopservice.SubmitLevelApplicationReq{
		ShopId:      c.ShopId,
		TargetLevel: req.TargetLevel,
	})
	if err != nil {
		return nil, err
	}
	return &types.SubmitLevelApplicationResp{ApplicationId: resp.ApplicationId}, nil
}

// ListLevelApplications returns paginated level upgrade applications for admin.
func ListLevelApplications(ctx context.Context, svcCtx *svc.ServiceContext, req *types.ListLevelApplicationsReq) (*types.ListLevelApplicationsResp, error) {
	resp, err := svcCtx.ShopRpc.ListLevelApplications(ctx, &shopservice.ListLevelApplicationsReq{
		Status:   req.Status,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	out := make([]*types.ShopLevelApplicationInfo, 0, len(resp.Applications))
	for _, a := range resp.Applications {
		out = append(out, mapLevelApplication(a))
	}
	return &types.ListLevelApplicationsResp{Total: resp.Total, Applications: out}, nil
}

// ReviewLevelApplication approves or rejects a pending level application.
func ReviewLevelApplication(ctx context.Context, svcCtx *svc.ServiceContext, id int64, req *types.ReviewLevelApplicationReq) (*types.OkResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	var adminId int64
	if c != nil {
		adminId = c.Uid
	}
	if _, err := svcCtx.ShopRpc.ReviewLevelApplication(ctx, &shopservice.ReviewLevelApplicationReq{
		ApplicationId: id,
		AdminId:       adminId,
		Action:        req.Action,
		Remark:        req.Remark,
	}); err != nil {
		return nil, err
	}
	return &types.OkResp{Ok: true}, nil
}

func mapLevelTemplate(t *shopservice.ShopLevelTemplate) *types.ShopLevelTemplateInfo {
	if t == nil {
		return nil
	}
	return &types.ShopLevelTemplateInfo{
		Level:          t.Level,
		Name:           t.Name,
		MinGmv:         t.MinGmv,
		MinCreditScore: t.MinCreditScore,
		MinMonths:      t.MinMonths,
		MinRating:      t.MinRating,
		CommissionRate: t.CommissionRate,
		TrafficBoost:   t.TrafficBoost,
		Benefits:       t.Benefits,
	}
}

func mapLevelApplication(a *shopservice.ShopLevelApplication) *types.ShopLevelApplicationInfo {
	return &types.ShopLevelApplicationInfo{
		Id:           a.Id,
		ShopId:       a.ShopId,
		CurrentLevel: a.CurrentLevel,
		TargetLevel:  a.TargetLevel,
		Snapshot:     a.Snapshot,
		Status:       a.Status,
		AdminId:      a.AdminId,
		AdminRemark:  a.AdminRemark,
		CreateTime:   a.CreateTime,
	}
}
