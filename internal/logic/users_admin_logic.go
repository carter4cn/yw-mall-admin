package logic

import (
	"context"

	"mall-admin-api/internal/svc"
	"mall-admin-api/internal/types"

	"mall-user-rpc/userclient"
)

func ListUsers(ctx context.Context, svcCtx *svc.ServiceContext, req *types.ListUsersReq) (*types.ListUsersResp, error) {
	resp, err := svcCtx.UserRpc.ListUsers(ctx, &userclient.ListUsersReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Keyword:  req.Keyword,
	})
	if err != nil {
		return nil, err
	}
	out := make([]*types.UserBrief, 0, len(resp.Users))
	for _, u := range resp.Users {
		out = append(out, &types.UserBrief{
			Id:         u.Id,
			Username:   u.Username,
			Phone:      u.Phone,
			Avatar:     u.Avatar,
			CreateTime: u.CreateTime,
		})
	}
	return &types.ListUsersResp{Total: resp.Total, Users: out}, nil
}

func UpdateUserStatus(ctx context.Context, svcCtx *svc.ServiceContext, id int64, req *types.UpdateUserStatusReq) (*types.OkResp, error) {
	if _, err := svcCtx.UserRpc.UpdateUserStatus(ctx, &userclient.UpdateUserStatusReq{
		Id:     id,
		Status: req.Status,
	}); err != nil {
		return nil, err
	}
	return &types.OkResp{Ok: true}, nil
}
