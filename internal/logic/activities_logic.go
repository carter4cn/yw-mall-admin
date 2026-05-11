package logic

import (
	"context"
	"errors"

	"mall-admin-api/internal/svc"
	"mall-admin-api/internal/types"

	"mall-activity-rpc/activityclient"
	"mall-rule-rpc/ruleclient"
)

func mapActivity(a *activityclient.ActivityInfo) *types.ActivityInfo {
	return &types.ActivityInfo{
		Id:                   a.Id,
		Code:                 a.Code,
		Title:                a.Title,
		Description:          a.Description,
		Type:                 a.Type,
		Status:               a.Status,
		StartTime:            a.StartTime,
		EndTime:              a.EndTime,
		RuleSetId:            a.RuleSetId,
		WorkflowDefinitionId: a.WorkflowDefinitionId,
		TemplateId:           a.TemplateId,
		ConfigJson:           a.ConfigJson,
		CreateTime:           a.CreateTime,
		UpdateTime:           a.UpdateTime,
	}
}

func AdminListActivities(ctx context.Context, svcCtx *svc.ServiceContext, req *types.AdminListActivitiesReq) (*types.ListActivitiesResp, error) {
	resp, err := svcCtx.ActivityRpc.ListActivities(ctx, &activityclient.ListActivitiesReq{
		Type:     req.Type,
		Status:   req.Status,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	out := make([]*types.ActivityInfo, 0, len(resp.Activities))
	for _, a := range resp.Activities {
		out = append(out, mapActivity(a))
	}
	return &types.ListActivitiesResp{Total: resp.Total, Activities: out}, nil
}

func MerchantListActivities(ctx context.Context, svcCtx *svc.ServiceContext, req *types.AdminListActivitiesReq) (*types.ListActivitiesResp, error) {
	// Merchant view currently mirrors admin list (no shop scoping in proto yet).
	return AdminListActivities(ctx, svcCtx, req)
}

func AdminCreateActivity(ctx context.Context, svcCtx *svc.ServiceContext, req *types.AdminCreateActivityReq) (*types.AdminCreateActivityResp, error) {
	resp, err := svcCtx.ActivityRpc.CreateActivity(ctx, &activityclient.CreateActivityReq{
		Code:                 req.Code,
		Title:                req.Title,
		Description:          req.Description,
		Type:                 req.Type,
		StartTime:            req.StartTime,
		EndTime:              req.EndTime,
		TemplateId:           req.TemplateId,
		RuleSetId:            req.RuleSetId,
		WorkflowDefinitionId: req.WorkflowDefinitionId,
		ConfigJson:           req.ConfigJson,
	})
	if err != nil {
		return nil, err
	}
	return &types.AdminCreateActivityResp{Id: resp.Id}, nil
}

func AdminUpdateActivity(ctx context.Context, svcCtx *svc.ServiceContext, id int64, req *types.AdminUpdateActivityReq) (*types.OkResp, error) {
	if _, err := svcCtx.ActivityRpc.UpdateActivity(ctx, &activityclient.UpdateActivityReq{
		Id:          id,
		Title:       req.Title,
		Description: req.Description,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		ConfigJson:  req.ConfigJson,
	}); err != nil {
		return nil, err
	}
	return &types.OkResp{Ok: true}, nil
}

func AdminSetActivityStatus(ctx context.Context, svcCtx *svc.ServiceContext, id int64, req *types.AdminSetActivityStatusReq) (*types.OkResp, error) {
	idReq := &activityclient.IdReq{Id: id}
	switch req.Status {
	case "PUBLISH", "PUBLISHED":
		if _, err := svcCtx.ActivityRpc.PublishActivity(ctx, idReq); err != nil {
			return nil, err
		}
	case "PAUSE", "PAUSED":
		if _, err := svcCtx.ActivityRpc.PauseActivity(ctx, idReq); err != nil {
			return nil, err
		}
	case "END", "ENDED":
		if _, err := svcCtx.ActivityRpc.EndActivity(ctx, idReq); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("invalid status; use PUBLISH/PAUSE/END")
	}
	return &types.OkResp{Ok: true}, nil
}

// ===== Rules =====

func ListRules(ctx context.Context, svcCtx *svc.ServiceContext, req *types.ListRulesReq) (*types.ListRulesResp, error) {
	resp, err := svcCtx.RuleRpc.ListRules(ctx, &ruleclient.ListRulesReq{Page: req.Page, PageSize: req.PageSize})
	if err != nil {
		return nil, err
	}
	out := make([]*types.RuleInfoBrief, 0, len(resp.Rules))
	for _, r := range resp.Rules {
		out = append(out, &types.RuleInfoBrief{
			Id:          r.Id,
			Code:        r.Code,
			Description: r.Description,
			Expression:  r.Expression,
			Status:      r.Status,
			CreateTime:  r.CreateTime,
		})
	}
	return &types.ListRulesResp{Total: resp.Total, Rules: out}, nil
}

func CreateRule(ctx context.Context, svcCtx *svc.ServiceContext, req *types.CreateRuleReq) (*types.CreateRuleResp, error) {
	resp, err := svcCtx.RuleRpc.CreateRule(ctx, &ruleclient.CreateRuleReq{
		Code:        req.Code,
		Description: req.Description,
		Expression:  req.Expression,
		JsonSchema:  req.JsonSchema,
	})
	if err != nil {
		return nil, err
	}
	return &types.CreateRuleResp{Id: resp.Id}, nil
}

func ValidateExpression(ctx context.Context, svcCtx *svc.ServiceContext, req *types.ValidateExpressionReq) (*types.ValidateExpressionResp, error) {
	resp, err := svcCtx.RuleRpc.ValidateExpression(ctx, &ruleclient.ValidateExpressionReq{
		Expression: req.Expression,
		Lang:       req.Lang,
	})
	if err != nil {
		return nil, err
	}
	return &types.ValidateExpressionResp{Valid: resp.Valid, Error: resp.Error}, nil
}

func CreateActivityRule(ctx context.Context, svcCtx *svc.ServiceContext, req *types.CreateActivityRuleReq) (*types.CreateActivityRuleResp, error) {
	conds := make([]*ruleclient.ActivityRuleCondition, 0, len(req.Conditions))
	for _, c := range req.Conditions {
		conds = append(conds, &ruleclient.ActivityRuleCondition{Type: c.Type, Operator: c.Operator, Value: c.Value, Unit: c.Unit})
	}
	excls := make([]*ruleclient.ActivityRuleExclusion, 0, len(req.Exclusions))
	for _, e := range req.Exclusions {
		excls = append(excls, &ruleclient.ActivityRuleExclusion{Type: e.Type, Value: e.Value})
	}
	rewards := make([]*ruleclient.ActivityRuleReward, 0, len(req.Rewards))
	for _, r := range req.Rewards {
		rewards = append(rewards, &ruleclient.ActivityRuleReward{Type: r.Type, Amount: r.Amount, Unit: r.Unit})
	}
	resp, err := svcCtx.RuleRpc.CreateActivityRule(ctx, &ruleclient.CreateActivityRuleReq{
		Code:        req.Code,
		Description: req.Description,
		Budget:      req.Budget,
		Conditions:  conds,
		Exclusions:  excls,
		Rewards:     rewards,
	})
	if err != nil {
		return nil, err
	}
	return &types.CreateActivityRuleResp{
		RuleId:              resp.RuleId,
		RuleSetId:           resp.RuleSetId,
		GeneratedExpression: resp.GeneratedExpression,
	}, nil
}
