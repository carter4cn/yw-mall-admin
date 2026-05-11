package logic

import (
	"context"
	"errors"
	"strings"

	"mall-admin-api/internal/middleware"
	"mall-admin-api/internal/svc"
	"mall-admin-api/internal/types"

	"mall-shop-rpc/shopservice"
	"mall-user-rpc/userclient"
)

// AdminLogin authenticates an admin and issues a JWT.
func AdminLogin(ctx context.Context, svcCtx *svc.ServiceContext, req *types.LoginReq) (*types.LoginResp, error) {
	resp, err := svcCtx.UserRpc.AdminLogin(ctx, &userclient.AdminLoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	perms := splitPerms(resp.Permissions)
	tok, err := middleware.IssueToken(resp.Id, "admin", 0, perms, svcCtx.JwtSecret.Get(), svcCtx.Config.Auth.AccessExpire)
	if err != nil {
		return nil, err
	}
	return &types.LoginResp{
		Token:       tok,
		Id:          resp.Id,
		Username:    resp.Username,
		Role:        "admin",
		Permissions: perms,
	}, nil
}

// MerchantLogin reuses the regular user login then attaches shop_id from shop service.
func MerchantLogin(ctx context.Context, svcCtx *svc.ServiceContext, req *types.LoginReq) (*types.LoginResp, error) {
	loginResp, err := svcCtx.UserRpc.Login(ctx, &userclient.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	shop, err := svcCtx.ShopRpc.GetShopByOwnerId(ctx, &shopservice.GetShopByOwnerIdReq{OwnerUserId: loginResp.Id})
	if err != nil {
		return nil, errors.New("merchant has no active shop; apply for one first")
	}
	tok, err := middleware.IssueToken(loginResp.Id, "merchant", shop.Id, nil, svcCtx.JwtSecret.Get(), svcCtx.Config.Auth.AccessExpire)
	if err != nil {
		return nil, err
	}
	return &types.LoginResp{
		Token:    tok,
		Id:       loginResp.Id,
		Username: req.Username,
		Role:     "merchant",
		ShopId:   shop.Id,
	}, nil
}

func CreateAdmin(ctx context.Context, svcCtx *svc.ServiceContext, req *types.CreateAdminReq) (*types.CreateAdminResp, error) {
	resp, err := svcCtx.UserRpc.CreateAdmin(ctx, &userclient.CreateAdminReq{
		Username:    req.Username,
		Password:    req.Password,
		Email:       req.Email,
		Role:        req.Role,
		Permissions: req.Permissions,
	})
	if err != nil {
		return nil, err
	}
	return &types.CreateAdminResp{Id: resp.Id}, nil
}

func ListAdmins(ctx context.Context, svcCtx *svc.ServiceContext, req *types.ListAdminsReq) (*types.ListAdminsResp, error) {
	resp, err := svcCtx.UserRpc.ListAdmins(ctx, &userclient.ListAdminsReq{
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	out := make([]*types.AdminInfo, 0, len(resp.Admins))
	for _, a := range resp.Admins {
		out = append(out, &types.AdminInfo{
			Id:          a.Id,
			Username:    a.Username,
			Email:       a.Email,
			Role:        a.Role,
			Permissions: a.Permissions,
			Status:      a.Status,
			CreateTime:  a.CreateTime,
		})
	}
	return &types.ListAdminsResp{Total: resp.Total, Admins: out}, nil
}

func splitPerms(s string) []string {
	if s == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}
