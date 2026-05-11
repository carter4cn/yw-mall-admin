package logic

import (
	"context"
	"errors"

	"mall-admin-api/internal/middleware"
	"mall-admin-api/internal/svc"
	"mall-admin-api/internal/types"

	"mall-payment-rpc/paymentclient"
)

func GetMerchantWallet(ctx context.Context, svcCtx *svc.ServiceContext) (*types.MerchantWalletInfo, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	if c == nil || c.ShopId <= 0 {
		return nil, errors.New("merchant shop unknown")
	}
	w, err := svcCtx.PaymentRpc.GetMerchantWallet(ctx, &paymentclient.GetMerchantWalletReq{ShopId: c.ShopId})
	if err != nil {
		return nil, err
	}
	return &types.MerchantWalletInfo{
		ShopId:         w.ShopId,
		Balance:        w.Balance,
		Frozen:         w.Frozen,
		TotalIncome:    w.TotalIncome,
		TotalWithdrawn: w.TotalWithdrawn,
		UpdateTime:     w.UpdateTime,
	}, nil
}

func ListBillRecords(ctx context.Context, svcCtx *svc.ServiceContext, req *types.ListBillRecordsReq) (*types.ListBillRecordsResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	if c == nil || c.ShopId <= 0 {
		return nil, errors.New("merchant shop unknown")
	}
	resp, err := svcCtx.PaymentRpc.ListBillRecords(ctx, &paymentclient.ListBillRecordsReq{
		ShopId:   c.ShopId,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	out := make([]*types.BillRecordInfo, 0, len(resp.Records))
	for _, r := range resp.Records {
		out = append(out, &types.BillRecordInfo{
			Id:         r.Id,
			ShopId:     r.ShopId,
			Type:       r.Type,
			Amount:     r.Amount,
			OrderId:    r.OrderId,
			Remark:     r.Remark,
			CreateTime: r.CreateTime,
		})
	}
	return &types.ListBillRecordsResp{Total: resp.Total, Records: out}, nil
}

func CreateWithdrawal(ctx context.Context, svcCtx *svc.ServiceContext, req *types.CreateWithdrawalReq) (*types.CreateWithdrawalResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	if c == nil || c.ShopId <= 0 {
		return nil, errors.New("merchant shop unknown")
	}
	resp, err := svcCtx.PaymentRpc.CreateWithdrawal(ctx, &paymentclient.CreateWithdrawalReq{
		ShopId:   c.ShopId,
		Amount:   req.Amount,
		BankInfo: req.BankInfo,
	})
	if err != nil {
		return nil, err
	}
	return &types.CreateWithdrawalResp{Id: resp.Id}, nil
}

func mapWithdrawals(rows []*paymentclient.WithdrawalInfo) []*types.WithdrawalInfo {
	out := make([]*types.WithdrawalInfo, 0, len(rows))
	for _, w := range rows {
		out = append(out, &types.WithdrawalInfo{
			Id:          w.Id,
			ShopId:      w.ShopId,
			Amount:      w.Amount,
			BankInfo:    w.BankInfo,
			Status:      w.Status,
			AdminId:     w.AdminId,
			AdminRemark: w.AdminRemark,
			CreateTime:  w.CreateTime,
			UpdateTime:  w.UpdateTime,
		})
	}
	return out
}

func ListMerchantWithdrawals(ctx context.Context, svcCtx *svc.ServiceContext, req *types.ListWithdrawalsReq) (*types.ListWithdrawalsResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	if c == nil || c.ShopId <= 0 {
		return nil, errors.New("merchant shop unknown")
	}
	resp, err := svcCtx.PaymentRpc.ListWithdrawals(ctx, &paymentclient.ListWithdrawalsReq{
		ShopId:   c.ShopId,
		Status:   req.Status,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	return &types.ListWithdrawalsResp{Total: resp.Total, Withdrawals: mapWithdrawals(resp.Withdrawals)}, nil
}

func AdminListWithdrawals(ctx context.Context, svcCtx *svc.ServiceContext, req *types.ListWithdrawalsReq) (*types.ListWithdrawalsResp, error) {
	resp, err := svcCtx.PaymentRpc.ListWithdrawals(ctx, &paymentclient.ListWithdrawalsReq{
		ShopId:   0,
		Status:   req.Status,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	return &types.ListWithdrawalsResp{Total: resp.Total, Withdrawals: mapWithdrawals(resp.Withdrawals)}, nil
}

func AdminHandleWithdrawal(ctx context.Context, svcCtx *svc.ServiceContext, id int64, req *types.AdminHandleWithdrawalReq) (*types.OkResp, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	var adminId int64
	if c != nil {
		adminId = c.Uid
	}
	if _, err := svcCtx.PaymentRpc.AdminHandleWithdrawal(ctx, &paymentclient.AdminHandleWithdrawalReq{
		Id:      id,
		AdminId: adminId,
		Action:  req.Action,
		Remark:  req.Remark,
	}); err != nil {
		return nil, err
	}
	return &types.OkResp{Ok: true}, nil
}
