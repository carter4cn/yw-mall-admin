package logic

import (
	"context"
	"errors"

	"mall-admin-api/internal/middleware"
	"mall-admin-api/internal/svc"
	"mall-admin-api/internal/types"

	"mall-logistics-rpc/logisticsclient"
)

// CreateFreightTemplate creates a new freight template for the current shop.
func CreateFreightTemplate(ctx context.Context, svcCtx *svc.ServiceContext, req *types.CreateFreightTemplateReq) (*types.CreateFreightTemplateResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	if c == nil || c.ShopId <= 0 {
		return nil, errors.New("merchant shop unknown")
	}
	resp, err := svcCtx.LogisticsRpc.CreateFreightTemplate(ctx, &logisticsclient.CreateFreightTemplateReq{
		ShopId:     c.ShopId,
		Name:       req.Name,
		CalcType:   req.CalcType,
		FirstValue: req.FirstValue,
		FirstFee:   req.FirstFee,
		ExtraValue: req.ExtraValue,
		ExtraFee:   req.ExtraFee,
		Regions:    req.Regions,
		IsDefault:  req.IsDefault,
	})
	if err != nil {
		return nil, err
	}
	return &types.CreateFreightTemplateResp{Id: resp.Id}, nil
}

func ListFreightTemplates(ctx context.Context, svcCtx *svc.ServiceContext, req *types.ListFreightTemplatesReq) (*types.ListFreightTemplatesResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	if c == nil || c.ShopId <= 0 {
		return nil, errors.New("merchant shop unknown")
	}
	resp, err := svcCtx.LogisticsRpc.ListFreightTemplates(ctx, &logisticsclient.ListFreightTemplatesReq{
		ShopId:   c.ShopId,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	out := make([]*types.FreightTemplateInfo, 0, len(resp.Templates))
	for _, t := range resp.Templates {
		out = append(out, mapFreightTemplate(t))
	}
	return &types.ListFreightTemplatesResp{Total: resp.Total, Templates: out}, nil
}

func GetFreightTemplate(ctx context.Context, svcCtx *svc.ServiceContext, id int64) (*types.FreightTemplateInfo, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	if c == nil || c.ShopId <= 0 {
		return nil, errors.New("merchant shop unknown")
	}
	t, err := svcCtx.LogisticsRpc.GetFreightTemplate(ctx, &logisticsclient.IdReq{Id: id})
	if err != nil {
		return nil, err
	}
	if t.ShopId != c.ShopId {
		return nil, errors.New("template not found")
	}
	return mapFreightTemplate(t), nil
}

func UpdateFreightTemplate(ctx context.Context, svcCtx *svc.ServiceContext, id int64, req *types.UpdateFreightTemplateReq) (*types.OkResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	if c == nil || c.ShopId <= 0 {
		return nil, errors.New("merchant shop unknown")
	}
	if _, err := svcCtx.LogisticsRpc.UpdateFreightTemplate(ctx, &logisticsclient.UpdateFreightTemplateReq{
		Id:        id,
		ShopId:    c.ShopId,
		Name:      req.Name,
		FirstFee:  req.FirstFee,
		ExtraFee:  req.ExtraFee,
		IsDefault: req.IsDefault,
	}); err != nil {
		return nil, err
	}
	return &types.OkResp{Ok: true}, nil
}

func DeleteFreightTemplate(ctx context.Context, svcCtx *svc.ServiceContext, id int64) (*types.OkResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	if c == nil || c.ShopId <= 0 {
		return nil, errors.New("merchant shop unknown")
	}
	// Verify ownership before delete.
	t, err := svcCtx.LogisticsRpc.GetFreightTemplate(ctx, &logisticsclient.IdReq{Id: id})
	if err != nil {
		return nil, err
	}
	if t.ShopId != c.ShopId {
		return nil, errors.New("template not found")
	}
	if _, err := svcCtx.LogisticsRpc.DeleteFreightTemplate(ctx, &logisticsclient.IdReq{Id: id}); err != nil {
		return nil, err
	}
	return &types.OkResp{Ok: true}, nil
}

func mapFreightTemplate(t *logisticsclient.FreightTemplate) *types.FreightTemplateInfo {
	return &types.FreightTemplateInfo{
		Id:         t.Id,
		ShopId:     t.ShopId,
		Name:       t.Name,
		CalcType:   t.CalcType,
		FirstValue: t.FirstValue,
		FirstFee:   t.FirstFee,
		ExtraValue: t.ExtraValue,
		ExtraFee:   t.ExtraFee,
		Regions:    t.Regions,
		IsDefault:  t.IsDefault,
		Status:     t.Status,
		CreateTime: t.CreateTime,
	}
}
