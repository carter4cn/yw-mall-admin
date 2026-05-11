package logic

import (
	"context"

	"mall-admin-api/internal/svc"
	"mall-admin-api/internal/types"

	"mall-risk-rpc/riskclient"
)

func CreateSensitiveWord(ctx context.Context, svcCtx *svc.ServiceContext, req *types.CreateSensitiveWordReq) (*types.CreateSensitiveWordResp, error) {
	resp, err := svcCtx.RiskRpc.CreateSensitiveWord(ctx, &riskclient.CreateSensitiveWordReq{
		Word:     req.Word,
		Category: req.Category,
		Action:   req.Action,
	})
	if err != nil {
		return nil, err
	}
	return &types.CreateSensitiveWordResp{Id: resp.Id}, nil
}

func ListSensitiveWords(ctx context.Context, svcCtx *svc.ServiceContext, req *types.ListSensitiveWordsReq) (*types.ListSensitiveWordsResp, error) {
	resp, err := svcCtx.RiskRpc.ListSensitiveWords(ctx, &riskclient.ListSensitiveWordsReq{
		Category: req.Category,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	out := make([]*types.SensitiveWordInfo, 0, len(resp.Words))
	for _, w := range resp.Words {
		out = append(out, &types.SensitiveWordInfo{
			Id:         w.Id,
			Word:       w.Word,
			Category:   w.Category,
			Action:     w.Action,
			Status:     w.Status,
			CreateTime: w.CreateTime,
		})
	}
	return &types.ListSensitiveWordsResp{Total: resp.Total, Words: out}, nil
}

func DeleteSensitiveWord(ctx context.Context, svcCtx *svc.ServiceContext, id int64) (*types.OkResp, error) {
	if _, err := svcCtx.RiskRpc.DeleteSensitiveWord(ctx, &riskclient.IdReq{Id: id}); err != nil {
		return nil, err
	}
	return &types.OkResp{Ok: true}, nil
}
