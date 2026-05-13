package logic

import (
	"context"

	"mall-admin-api/internal/svc"
	"mall-admin-api/internal/types"

	"mall-payment-rpc/paymentclient"
)

// ListLedger paginates account_ledger rows; admin-only, accepts optional
// shop_id / category / time-window filters.
func ListLedger(ctx context.Context, svcCtx *svc.ServiceContext, req *types.ListLedgerReq) (*types.ListLedgerResp, error) {
	resp, err := svcCtx.PaymentRpc.ListLedger(ctx, &paymentclient.ListLedgerReq{
		ShopId:    req.ShopId,
		Category:  req.Category,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Page:      req.Page,
		PageSize:  req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	entries := make([]types.LedgerEntryDTO, 0, len(resp.Entries))
	for _, e := range resp.Entries {
		entries = append(entries, types.LedgerEntryDTO{
			Id:             e.Id,
			ShopId:         e.ShopId,
			Direction:      e.Direction,
			Category:       e.Category,
			Amount:         e.Amount,
			RunningBalance: e.RunningBalance,
			OrderId:        e.OrderId,
			RefundId:       e.RefundId,
			RefNo:          e.RefNo,
			Description:    e.Description,
			CreateTime:     e.CreateTime,
		})
	}
	return &types.ListLedgerResp{Entries: entries, Total: resp.Total}, nil
}

func GetLedgerSummary(ctx context.Context, svcCtx *svc.ServiceContext, req *types.GetLedgerSummaryReq) (*types.LedgerSummaryDTO, error) {
	resp, err := svcCtx.PaymentRpc.GetShopLedgerSummary(ctx, &paymentclient.GetShopLedgerSummaryReq{
		ShopId:    req.ShopId,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	})
	if err != nil {
		return nil, err
	}
	return &types.LedgerSummaryDTO{
		TotalIncome:     resp.TotalIncome,
		TotalRefund:     resp.TotalRefund,
		TotalCommission: resp.TotalCommission,
		TotalWithdrawal: resp.TotalWithdrawal,
		NetBalance:      resp.NetBalance,
	}, nil
}

func RunReconcile(ctx context.Context, svcCtx *svc.ServiceContext, req *types.RunReconcileReq) (*types.ReconcileReportDTO, error) {
	resp, err := svcCtx.PaymentRpc.RunReconciliation(ctx, &paymentclient.RunReconciliationReq{
		ShopId: req.ShopId,
	})
	if err != nil {
		return nil, err
	}
	results := make([]types.ShopReconcileResultDTO, 0, len(resp.Results))
	for _, r := range resp.Results {
		results = append(results, types.ShopReconcileResultDTO{
			ShopId:        r.ShopId,
			LedgerCredit:  r.LedgerCredit,
			LedgerDebit:   r.LedgerDebit,
			LedgerNet:     r.LedgerNet,
			WalletBalance: r.WalletBalance,
			WalletFrozen:  r.WalletFrozen,
			WalletTotal:   r.WalletTotal,
			Diff:          r.Diff,
			Passed:        r.Passed,
		})
	}
	return &types.ReconcileReportDTO{
		TotalChecked: resp.TotalChecked,
		Passed:       resp.Passed,
		Failed:       resp.Failed,
		Results:      results,
	}, nil
}
