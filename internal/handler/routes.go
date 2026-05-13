package handler

import (
	"net/http"

	"mall-admin-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

// RegisterHandlers wires every admin/merchant route, applying the matching
// auth + op-log middleware per route group.
func RegisterHandlers(server *rest.Server, svcCtx *svc.ServiceContext) {
	// ===== /admin/v1 (public sub-routes) =====
	server.AddRoutes(
		[]rest.Route{
			{Method: http.MethodPost, Path: "/login", Handler: adminLoginHandler(svcCtx)},
		},
		rest.WithPrefix("/admin/v1"),
	)

	// ===== /admin/v1 (protected) =====
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{svcCtx.AdminAuth, svcCtx.OpLog},
			[]rest.Route{
				{Method: http.MethodPost, Path: "/accounts", Handler: createAdminHandler(svcCtx)},
				{Method: http.MethodGet, Path: "/accounts", Handler: listAdminsHandler(svcCtx)},

				{Method: http.MethodGet, Path: "/shop-applications", Handler: listShopApplicationsHandler(svcCtx)},
				{Method: http.MethodGet, Path: "/shop-applications/:id", Handler: getShopApplicationHandler(svcCtx)},
				{Method: http.MethodPost, Path: "/shop-applications/:id/review", Handler: reviewShopApplicationHandler(svcCtx)},

				{Method: http.MethodGet, Path: "/shops", Handler: listShopsHandler(svcCtx)},
				{Method: http.MethodPost, Path: "/shops/:id/status", Handler: updateShopStatusHandler(svcCtx)},
				{Method: http.MethodPost, Path: "/shops/:id/credit", Handler: adjustCreditScoreHandler(svcCtx)},

				{Method: http.MethodGet, Path: "/products/review", Handler: adminListReviewProductsHandler(svcCtx)},
				{Method: http.MethodPost, Path: "/products/:id/review", Handler: adminReviewProductHandler(svcCtx)},

				{Method: http.MethodGet, Path: "/users", Handler: listUsersHandler(svcCtx)},
				{Method: http.MethodPost, Path: "/users/:id/status", Handler: updateUserStatusHandler(svcCtx)},

				// ----- P1 Epic E: review takedown -----
				{Method: http.MethodGet, Path: "/reviews/delete-requests", Handler: listDeleteRequestsHandler(svcCtx)},
				{Method: http.MethodPost, Path: "/reviews/delete-requests/:id/handle", Handler: adminHandleDeleteRequestHandler(svcCtx)},

				// ----- P1 Epic F: complaints & restrictions -----
				{Method: http.MethodGet, Path: "/complaints", Handler: listComplaintsHandler(svcCtx)},
				{Method: http.MethodGet, Path: "/complaints/:id", Handler: getComplaintHandler(svcCtx)},
				{Method: http.MethodPost, Path: "/complaints/:id/handle", Handler: handleComplaintHandler(svcCtx)},
				{Method: http.MethodPost, Path: "/shops/:id/restrictions", Handler: setShopRestrictionHandler(svcCtx)},
				{Method: http.MethodGet, Path: "/shops/:id/restrictions", Handler: listShopRestrictionsHandler(svcCtx)},
				{Method: http.MethodDelete, Path: "/shops/:id/restrictions/:rid", Handler: removeShopRestrictionHandler(svcCtx)},

				// ----- P1 Epic G: activities & rules -----
				{Method: http.MethodGet, Path: "/activities", Handler: adminListActivitiesHandler(svcCtx)},
				{Method: http.MethodPost, Path: "/activities", Handler: adminCreateActivityHandler(svcCtx)},
				{Method: http.MethodPut, Path: "/activities/:id", Handler: adminUpdateActivityHandler(svcCtx)},
				{Method: http.MethodPost, Path: "/activities/:id/status", Handler: adminSetActivityStatusHandler(svcCtx)},
				{Method: http.MethodGet, Path: "/rules", Handler: listRulesHandler(svcCtx)},
				{Method: http.MethodPost, Path: "/rules", Handler: createRuleHandler(svcCtx)},
				{Method: http.MethodPost, Path: "/rules/validate", Handler: validateExpressionHandler(svcCtx)},
				{Method: http.MethodPost, Path: "/activity-rules", Handler: createActivityRuleHandler(svcCtx)},

				// ----- P1 Epic H: withdrawals (admin) -----
				{Method: http.MethodGet, Path: "/withdrawals", Handler: adminListWithdrawalsHandler(svcCtx)},
				{Method: http.MethodPost, Path: "/withdrawals/:id/handle", Handler: adminHandleWithdrawalHandler(svcCtx)},

				// ----- P2 B-5/B-6: shop level applications -----
				{Method: http.MethodGet, Path: "/level-applications", Handler: listLevelApplicationsHandler(svcCtx)},
				{Method: http.MethodPost, Path: "/level-applications/:id/review", Handler: reviewLevelApplicationHandler(svcCtx)},

				// ----- P2 B-4: shop lifecycle (deactivate/pause/resume) -----
				{Method: http.MethodGet, Path: "/shop-lifecycle-requests", Handler: listShopLifecycleRequestsHandler(svcCtx)},
				{Method: http.MethodPost, Path: "/shop-lifecycle-requests/:id/review", Handler: reviewShopLifecycleRequestHandler(svcCtx)},

				// ----- P2 J-4: sensitive words -----
				{Method: http.MethodPost, Path: "/sensitive-words", Handler: createSensitiveWordHandler(svcCtx)},
				{Method: http.MethodGet, Path: "/sensitive-words", Handler: listSensitiveWordsHandler(svcCtx)},
				{Method: http.MethodDelete, Path: "/sensitive-words/:id", Handler: deleteSensitiveWordHandler(svcCtx)},

				// ----- S2: refund arbitration -----
				{Method: http.MethodGet, Path: "/refunds/arbitrations", Handler: listPendingArbitrationsHandler(svcCtx)},
				{Method: http.MethodPost, Path: "/refunds/:id/arbitrate", Handler: arbitrateRefundHandler(svcCtx)},
			}...,
		),
		rest.WithPrefix("/admin/v1"),
	)

	// ===== /merchant/v1 (public sub-routes) =====
	server.AddRoutes(
		[]rest.Route{
			{Method: http.MethodPost, Path: "/login", Handler: merchantLoginHandler(svcCtx)},
			{Method: http.MethodPost, Path: "/apply", Handler: applyShopHandler(svcCtx)},
			{Method: http.MethodGet, Path: "/apply/:id", Handler: getMyApplicationHandler(svcCtx)},
		},
		rest.WithPrefix("/merchant/v1"),
	)

	// ===== /merchant/v1 (protected) =====
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{svcCtx.MerchantAuth, svcCtx.OpLog},
			[]rest.Route{
				{Method: http.MethodGet, Path: "/shop", Handler: getMyShopHandler(svcCtx)},
				{Method: http.MethodPut, Path: "/shop", Handler: updateMyShopHandler(svcCtx)},

				{Method: http.MethodGet, Path: "/products", Handler: merchantListProductsHandler(svcCtx)},
				{Method: http.MethodPost, Path: "/products", Handler: merchantCreateProductHandler(svcCtx)},
				{Method: http.MethodGet, Path: "/products/:id", Handler: merchantGetProductHandler(svcCtx)},
				{Method: http.MethodPut, Path: "/products/:id", Handler: merchantUpdateProductHandler(svcCtx)},
				{Method: http.MethodPost, Path: "/products/:id/status", Handler: setProductStatusHandler(svcCtx)},
				{Method: http.MethodPost, Path: "/products/:id/stock", Handler: setProductStockHandler(svcCtx)},

				{Method: http.MethodGet, Path: "/orders", Handler: merchantListOrdersHandler(svcCtx)},
				{Method: http.MethodGet, Path: "/orders/:id", Handler: merchantGetOrderHandler(svcCtx)},
				{Method: http.MethodPost, Path: "/orders/:id/ship", Handler: shipOrderHandler(svcCtx)},
				{Method: http.MethodPost, Path: "/orders/:id/reject-refund", Handler: merchantRejectRefundHandler(svcCtx)},

				// ----- P1 Epic E: reviews -----
				{Method: http.MethodGet, Path: "/reviews", Handler: merchantListReviewsHandler(svcCtx)},
				{Method: http.MethodPost, Path: "/reviews/:id/delete-request", Handler: requestDeleteReviewHandler(svcCtx)},

				// ----- P1 Epic F: complaints (merchant submits) -----
				{Method: http.MethodPost, Path: "/complaints", Handler: createComplaintHandler(svcCtx)},

				// ----- P1 Epic G: activities (merchant browses) -----
				{Method: http.MethodGet, Path: "/activities", Handler: merchantListActivitiesHandler(svcCtx)},

				// ----- P1 Epic H: wallet -----
				{Method: http.MethodGet, Path: "/wallet", Handler: getMerchantWalletHandler(svcCtx)},
				{Method: http.MethodGet, Path: "/wallet/bills", Handler: listBillRecordsHandler(svcCtx)},
				{Method: http.MethodPost, Path: "/wallet/withdraw", Handler: createWithdrawalHandler(svcCtx)},
				{Method: http.MethodGet, Path: "/wallet/withdrawals", Handler: listMerchantWithdrawalsHandler(svcCtx)},

				// ----- P2 B-5/B-6: shop levels -----
				{Method: http.MethodGet, Path: "/shop-levels", Handler: listShopLevelsHandler(svcCtx)},
				{Method: http.MethodGet, Path: "/shop/level-status", Handler: getMyLevelStatusHandler(svcCtx)},
				{Method: http.MethodPost, Path: "/shop/level/apply", Handler: submitLevelApplicationHandler(svcCtx)},

				// ----- P2 G-3 shop coupons -----
				{Method: http.MethodPost, Path: "/coupons", Handler: createShopCouponHandler(svcCtx)},
				{Method: http.MethodGet, Path: "/coupons", Handler: listShopCouponsHandler(svcCtx)},
				{Method: http.MethodPost, Path: "/coupons/:id/status", Handler: updateShopCouponStatusHandler(svcCtx)},

				// ----- P2 G-4 flash discounts -----
				{Method: http.MethodPost, Path: "/flash-discounts", Handler: createFlashDiscountHandler(svcCtx)},
				{Method: http.MethodGet, Path: "/flash-discounts", Handler: listFlashDiscountsHandler(svcCtx)},
				{Method: http.MethodPost, Path: "/flash-discounts/:id/cancel", Handler: cancelFlashDiscountHandler(svcCtx)},

				// ----- P2 B-4: shop lifecycle submit -----
				{Method: http.MethodPost, Path: "/shop/lifecycle", Handler: submitShopLifecycleRequestHandler(svcCtx)},

				// ----- P2 D-3: batch ship orders (CSV) -----
				{Method: http.MethodPost, Path: "/orders/batch-ship", Handler: batchShipOrdersHandler(svcCtx)},

				// ----- P2 C-6: freight templates -----
				{Method: http.MethodPost, Path: "/freight-templates", Handler: createFreightTemplateHandler(svcCtx)},
				{Method: http.MethodGet, Path: "/freight-templates", Handler: listFreightTemplatesHandler(svcCtx)},
				{Method: http.MethodGet, Path: "/freight-templates/:id", Handler: getFreightTemplateHandler(svcCtx)},
				{Method: http.MethodPut, Path: "/freight-templates/:id", Handler: updateFreightTemplateHandler(svcCtx)},
				{Method: http.MethodDelete, Path: "/freight-templates/:id", Handler: deleteFreightTemplateHandler(svcCtx)},

				// ----- S2: refund (merchant) -----
				{Method: http.MethodGet, Path: "/refunds", Handler: listShopRefundsHandler(svcCtx)},
				{Method: http.MethodPost, Path: "/refunds/:id/handle", Handler: merchantHandleRefundHandler(svcCtx)},
			}...,
		),
		rest.WithPrefix("/merchant/v1"),
	)
}
