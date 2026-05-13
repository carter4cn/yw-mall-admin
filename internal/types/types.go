package types

// ===== Auth =====
type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	Token       string   `json:"token"`
	Id          int64    `json:"id"`
	Username    string   `json:"username"`
	Role        string   `json:"role"`
	ShopId      int64    `json:"shopId,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
}

type CreateAdminReq struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email,omitempty"`
	Role        string `json:"role,omitempty"`
	Permissions string `json:"permissions,omitempty"`
}

type CreateAdminResp struct {
	Id int64 `json:"id"`
}

type ListAdminsReq struct {
	Page     int32 `form:"page,default=1"`
	PageSize int32 `form:"pageSize,default=20"`
}

type AdminInfo struct {
	Id          int64  `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	Permissions string `json:"permissions"`
	Status      int32  `json:"status"`
	CreateTime  int64  `json:"createTime"`
}

type ListAdminsResp struct {
	Total  int64        `json:"total"`
	Admins []*AdminInfo `json:"admins"`
}

// ===== Shop application & Shop admin =====
type ApplyShopReq struct {
	UserId          int64  `json:"userId"`
	ShopName        string `json:"shopName"`
	Logo            string `json:"logo"`
	Description     string `json:"description"`
	ContactPhone    string `json:"contactPhone"`
	BusinessLicense string `json:"businessLicense"`
	LegalPerson     string `json:"legalPerson"`
	IdCardFront     string `json:"idCardFront"`
	IdCardBack      string `json:"idCardBack"`
	Category        string `json:"category"`
}

type ApplyShopResp struct {
	ApplicationId int64 `json:"applicationId"`
}

type ShopApplication struct {
	Id              int64  `json:"id"`
	UserId          int64  `json:"userId"`
	ShopName        string `json:"shopName"`
	Logo            string `json:"logo"`
	Description     string `json:"description"`
	ContactPhone    string `json:"contactPhone"`
	BusinessLicense string `json:"businessLicense"`
	LegalPerson     string `json:"legalPerson"`
	IdCardFront     string `json:"idCardFront"`
	IdCardBack      string `json:"idCardBack"`
	Category        string `json:"category"`
	Status          int32  `json:"status"`
	ReviewRemark    string `json:"reviewRemark"`
	ReviewerId      int64  `json:"reviewerId"`
	ShopId          int64  `json:"shopId"`
	CreateTime      int64  `json:"createTime"`
	UpdateTime      int64  `json:"updateTime"`
}

type ListShopApplicationsReq struct {
	Status   int32 `form:"status,default=-1"`
	Page     int32 `form:"page,default=1"`
	PageSize int32 `form:"pageSize,default=20"`
}

type ListShopApplicationsResp struct {
	Total        int64              `json:"total"`
	Applications []*ShopApplication `json:"applications"`
}

type ReviewShopApplicationReq struct {
	Action int32  `json:"action"`
	Remark string `json:"remark,omitempty"`
}

type ListShopsReq struct {
	Page     int32 `form:"page,default=1"`
	PageSize int32 `form:"pageSize,default=20"`
}

type ShopBrief struct {
	Id           int64   `json:"id"`
	Name         string  `json:"name"`
	Logo         string  `json:"logo"`
	Status       int32   `json:"status"`
	CreateTime   int64   `json:"createTime"`
	Rating       float64 `json:"rating"`
	ProductCount int32   `json:"productCount"`
}

type ListShopsResp struct {
	Total int64        `json:"total"`
	Shops []*ShopBrief `json:"shops"`
}

type UpdateShopStatusReq struct {
	Status int32  `json:"status"`
	Reason string `json:"reason,omitempty"`
}

type AdjustCreditScoreReq struct {
	Delta  int32  `json:"delta"`
	Reason string `json:"reason,omitempty"`
}

// ===== Users =====
type ListUsersReq struct {
	Page     int32  `form:"page,default=1"`
	PageSize int32  `form:"pageSize,default=20"`
	Keyword  string `form:"keyword,optional"`
}

type UserBrief struct {
	Id         int64  `json:"id"`
	Username   string `json:"username"`
	Phone      string `json:"phone"`
	Avatar     string `json:"avatar"`
	CreateTime int64  `json:"createTime"`
}

type ListUsersResp struct {
	Total int64        `json:"total"`
	Users []*UserBrief `json:"users"`
}

type UpdateUserStatusReq struct {
	Status int32 `json:"status"`
}

// ===== Products =====
type AdminListReviewProductsReq struct {
	ReviewStatus int32 `form:"reviewStatus,default=0"`
	Page         int32 `form:"page,default=1"`
	PageSize     int32 `form:"pageSize,default=20"`
}

type ProductBrief struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Stock       int64  `json:"stock"`
	Images      string `json:"images"`
	ShopId      int64  `json:"shopId"`
	Status      int32  `json:"status"`
	CategoryId  int64  `json:"categoryId"`
	CreateTime  int64  `json:"createTime"`
}

type ListProductsResp struct {
	Total    int64           `json:"total"`
	Products []*ProductBrief `json:"products"`
}

type AdminReviewProductReq struct {
	Action int32  `json:"action"`
	Remark string `json:"remark,omitempty"`
}

type MerchantListProductsReq struct {
	Status       int32 `form:"status,default=-1"`
	ReviewStatus int32 `form:"reviewStatus,default=-1"`
	Page         int32 `form:"page,default=1"`
	PageSize     int32 `form:"pageSize,default=20"`
}

type MerchantCreateProductReq struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       int64   `json:"price"`
	Stock       int64   `json:"stock"`
	CategoryId  int64   `json:"categoryId"`
	Images      string  `json:"images"`
	Detail      string  `json:"detail,omitempty"`
	Brand       string  `json:"brand,omitempty"`
	Weight      float64 `json:"weight,omitempty"`
}

type MerchantCreateProductResp struct {
	Id int64 `json:"id"`
}

type MerchantUpdateProductReq struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       int64   `json:"price"`
	CategoryId  int64   `json:"categoryId"`
	Images      string  `json:"images"`
	Detail      string  `json:"detail,omitempty"`
	Brand       string  `json:"brand,omitempty"`
	Weight      float64 `json:"weight,omitempty"`
}

type SetProductStatusReq struct {
	Status int32 `json:"status"`
}

type SetProductStockReq struct {
	Stock int64 `json:"stock"`
}

type ProductDetail struct {
	Id           int64   `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        int64   `json:"price"`
	Stock        int64   `json:"stock"`
	CategoryId   int64   `json:"categoryId"`
	Images       string  `json:"images"`
	Status       int32   `json:"status"`
	ShopId       int64   `json:"shopId"`
	ReviewStatus int32   `json:"reviewStatus"`
	ReviewRemark string  `json:"reviewRemark"`
	Detail       string  `json:"detail"`
	Brand        string  `json:"brand"`
	Weight       float64 `json:"weight"`
	CreateTime   int64   `json:"createTime"`
}

// ===== Orders =====
type MerchantListOrdersReq struct {
	Status   int32 `form:"status,default=-1"`
	Page     int32 `form:"page,default=1"`
	PageSize int32 `form:"pageSize,default=20"`
}

type OrderItem struct {
	ProductId   int64  `json:"productId"`
	ProductName string `json:"productName"`
	Price       int64  `json:"price"`
	Quantity    int32  `json:"quantity"`
}

type OrderBrief struct {
	Id          int64        `json:"id"`
	OrderNo     string       `json:"orderNo"`
	UserId      int64        `json:"userId"`
	TotalAmount int64        `json:"totalAmount"`
	Status      int32        `json:"status"`
	CreateTime  int64        `json:"createTime"`
	Items       []*OrderItem `json:"items"`
}

type MerchantListOrdersResp struct {
	Total  int64         `json:"total"`
	Orders []*OrderBrief `json:"orders"`
}

type ShipOrderReq struct {
	Carrier    string `json:"carrier"`
	TrackingNo string `json:"trackingNo"`
}

type RejectRefundReq struct {
	Reason string `json:"reason"`
}

// ===== Generic =====
type OkResp struct {
	Ok bool `json:"ok"`
}

// ===== Shop =====
type ShopDetail struct {
	Id              int64   `json:"id"`
	Name            string  `json:"name"`
	Logo            string  `json:"logo"`
	Banner          string  `json:"banner"`
	Description     string  `json:"description"`
	Rating          float64 `json:"rating"`
	ProductCount    int32   `json:"productCount"`
	FollowCount     int32   `json:"followCount"`
	Status          int32   `json:"status"`
	CreateTime      int64   `json:"createTime"`
	OwnerUserId     int64   `json:"ownerUserId"`
	CreditScore     int32   `json:"creditScore"`
	Level           int32   `json:"level"`
	ContactPhone    string  `json:"contactPhone"`
	BusinessLicense string  `json:"businessLicense"`
}

type UpdateMyShopReq struct {
	Logo        string `json:"logo,omitempty"`
	Banner      string `json:"banner,omitempty"`
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
}

// ===== P1 Epic E: Reviews =====
type MerchantListReviewsReq struct {
	Score    int32 `form:"score,default=0"`
	Page     int32 `form:"page,default=1"`
	PageSize int32 `form:"pageSize,default=20"`
}

type ReviewBrief struct {
	Id                int64  `json:"id"`
	OrderItemId       int64  `json:"orderItemId"`
	UserId            int64  `json:"userId"`
	ProductId         int64  `json:"productId"`
	ScoreOverall      int32  `json:"scoreOverall"`
	Content           string `json:"content"`
	MerchantReplyText string `json:"merchantReplyText"`
	Status            int32  `json:"status"`
	CreateTime        int64  `json:"createTime"`
}

type ListReviewsResp struct {
	Total   int64          `json:"total"`
	Reviews []*ReviewBrief `json:"reviews"`
}

type RequestDeleteReviewReq struct {
	Reason string `json:"reason"`
}

type ListDeleteRequestsReq struct {
	Status   int32 `form:"status,default=-1"`
	Page     int32 `form:"page,default=1"`
	PageSize int32 `form:"pageSize,default=20"`
}

type ReviewDeleteRequestInfo struct {
	Id          int64  `json:"id"`
	ReviewId    int64  `json:"reviewId"`
	ShopId      int64  `json:"shopId"`
	Reason      string `json:"reason"`
	Status      int32  `json:"status"`
	AdminRemark string `json:"adminRemark"`
	AdminId     int64  `json:"adminId"`
	CreateTime  int64  `json:"createTime"`
}

type ListDeleteRequestsResp struct {
	Total    int64                      `json:"total"`
	Requests []*ReviewDeleteRequestInfo `json:"requests"`
}

type AdminHandleDeleteRequestReq struct {
	Action int32  `json:"action"`
	Remark string `json:"remark,omitempty"`
}

// ===== P1 Epic F: Complaints & risk =====
type CreateComplaintReq struct {
	DefendantType string `json:"defendantType"`
	DefendantId   int64  `json:"defendantId"`
	OrderId       int64  `json:"orderId,omitempty"`
	Category      string `json:"category"`
	Content       string `json:"content"`
	EvidenceUrls  string `json:"evidenceUrls,omitempty"`
}

type CreateComplaintResp struct {
	Id int64 `json:"id"`
}

type ListComplaintsReq struct {
	Status        int32  `form:"status,default=-1"`
	DefendantType string `form:"defendantType,optional"`
	DefendantId   int64  `form:"defendantId,optional"`
	Page          int32  `form:"page,default=1"`
	PageSize      int32  `form:"pageSize,default=20"`
}

type ComplaintTicketInfo struct {
	Id              int64  `json:"id"`
	ComplainantType string `json:"complainantType"`
	ComplainantId   int64  `json:"complainantId"`
	DefendantType   string `json:"defendantType"`
	DefendantId     int64  `json:"defendantId"`
	OrderId         int64  `json:"orderId"`
	Category        string `json:"category"`
	Content         string `json:"content"`
	EvidenceUrls    string `json:"evidenceUrls"`
	Status          int32  `json:"status"`
	AdminId         int64  `json:"adminId"`
	AdminRemark     string `json:"adminRemark"`
	CreateTime      int64  `json:"createTime"`
	UpdateTime      int64  `json:"updateTime"`
}

type ListComplaintsResp struct {
	Total   int64                  `json:"total"`
	Tickets []*ComplaintTicketInfo `json:"tickets"`
}

type HandleComplaintReq struct {
	Action int32  `json:"action"`
	Remark string `json:"remark,omitempty"`
}

type SetShopRestrictionReq struct {
	Restriction string `json:"restriction"`
	Reason      string `json:"reason,omitempty"`
	ExpireTime  int64  `json:"expireTime,omitempty"`
}

type ShopRestrictionInfo struct {
	Id          int64  `json:"id"`
	ShopId      int64  `json:"shopId"`
	Restriction string `json:"restriction"`
	Reason      string `json:"reason"`
	OperatorId  int64  `json:"operatorId"`
	ExpireTime  int64  `json:"expireTime"`
	CreateTime  int64  `json:"createTime"`
}

type ListShopRestrictionsResp struct {
	Restrictions []*ShopRestrictionInfo `json:"restrictions"`
}

// ===== P1 Epic G: Activities & Rules =====
type AdminListActivitiesReq struct {
	Type     string `form:"type,optional"`
	Status   string `form:"status,optional"`
	Page     int32  `form:"page,default=1"`
	PageSize int32  `form:"pageSize,default=20"`
}

type ActivityInfo struct {
	Id                   int64  `json:"id"`
	Code                 string `json:"code"`
	Title                string `json:"title"`
	Description          string `json:"description"`
	Type                 string `json:"type"`
	Status               string `json:"status"`
	StartTime            int64  `json:"startTime"`
	EndTime              int64  `json:"endTime"`
	RuleSetId            int64  `json:"ruleSetId"`
	WorkflowDefinitionId int64  `json:"workflowDefinitionId"`
	TemplateId           int64  `json:"templateId"`
	ConfigJson           string `json:"configJson"`
	CreateTime           int64  `json:"createTime"`
	UpdateTime           int64  `json:"updateTime"`
}

type ListActivitiesResp struct {
	Total      int64           `json:"total"`
	Activities []*ActivityInfo `json:"activities"`
}

type AdminCreateActivityReq struct {
	Code                 string `json:"code"`
	Title                string `json:"title"`
	Description          string `json:"description,omitempty"`
	Type                 string `json:"type"`
	StartTime            int64  `json:"startTime"`
	EndTime              int64  `json:"endTime"`
	TemplateId           int64  `json:"templateId,omitempty"`
	RuleSetId            int64  `json:"ruleSetId,omitempty"`
	WorkflowDefinitionId int64  `json:"workflowDefinitionId,omitempty"`
	ConfigJson           string `json:"configJson,omitempty"`
}

type AdminCreateActivityResp struct {
	Id int64 `json:"id"`
}

type AdminUpdateActivityReq struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	StartTime   int64  `json:"startTime,omitempty"`
	EndTime     int64  `json:"endTime,omitempty"`
	ConfigJson  string `json:"configJson,omitempty"`
}

type AdminSetActivityStatusReq struct {
	Status string `json:"status"` // PUBLISH / PAUSE / END
}

type ListRulesReq struct {
	Page     int32 `form:"page,default=1"`
	PageSize int32 `form:"pageSize,default=20"`
}

type RuleInfoBrief struct {
	Id          int64  `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Expression  string `json:"expression"`
	Status      string `json:"status"`
	CreateTime  int64  `json:"createTime"`
}

type ListRulesResp struct {
	Total int64            `json:"total"`
	Rules []*RuleInfoBrief `json:"rules"`
}

type CreateRuleReq struct {
	Code        string `json:"code"`
	Description string `json:"description,omitempty"`
	Expression  string `json:"expression"`
	JsonSchema  string `json:"jsonSchema,omitempty"`
}

type CreateRuleResp struct {
	Id int64 `json:"id"`
}

type ValidateExpressionReq struct {
	Expression string `json:"expression"`
	Lang       string `json:"lang,omitempty"`
}

type ValidateExpressionResp struct {
	Valid bool   `json:"valid"`
	Error string `json:"error,omitempty"`
}

type ActivityRuleConditionInput struct {
	Type     string `json:"type"`
	Operator string `json:"operator"`
	Value    int64  `json:"value"`
	Unit     string `json:"unit,omitempty"`
}

type ActivityRuleExclusionInput struct {
	Type  string `json:"type"`
	Value string `json:"value,omitempty"`
}

type ActivityRuleRewardInput struct {
	Type   string `json:"type"`
	Amount int64  `json:"amount"`
	Unit   string `json:"unit,omitempty"`
}

type CreateActivityRuleReq struct {
	Code        string                        `json:"code"`
	Description string                        `json:"description,omitempty"`
	Budget      int64                         `json:"budget,omitempty"`
	Conditions  []*ActivityRuleConditionInput `json:"conditions,omitempty"`
	Exclusions  []*ActivityRuleExclusionInput `json:"exclusions,omitempty"`
	Rewards     []*ActivityRuleRewardInput    `json:"rewards,omitempty"`
}

type CreateActivityRuleResp struct {
	RuleId              int64  `json:"ruleId"`
	RuleSetId           int64  `json:"ruleSetId"`
	GeneratedExpression string `json:"generatedExpression"`
}

// ===== P1 Epic H: Wallet & Withdrawals =====
type MerchantWalletInfo struct {
	ShopId         int64 `json:"shopId"`
	Balance        int64 `json:"balance"`
	Frozen         int64 `json:"frozen"`
	TotalIncome    int64 `json:"totalIncome"`
	TotalWithdrawn int64 `json:"totalWithdrawn"`
	UpdateTime     int64 `json:"updateTime"`
}

type ListBillRecordsReq struct {
	Page     int32 `form:"page,default=1"`
	PageSize int32 `form:"pageSize,default=20"`
}

type BillRecordInfo struct {
	Id         int64  `json:"id"`
	ShopId     int64  `json:"shopId"`
	Type       string `json:"type"`
	Amount     int64  `json:"amount"`
	OrderId    int64  `json:"orderId"`
	Remark     string `json:"remark"`
	CreateTime int64  `json:"createTime"`
}

type ListBillRecordsResp struct {
	Total   int64             `json:"total"`
	Records []*BillRecordInfo `json:"records"`
}

type CreateWithdrawalReq struct {
	Amount   int64  `json:"amount"`
	BankInfo string `json:"bankInfo"`
}

type CreateWithdrawalResp struct {
	Id int64 `json:"id"`
}

type ListWithdrawalsReq struct {
	Status   int32 `form:"status,default=-1"`
	Page     int32 `form:"page,default=1"`
	PageSize int32 `form:"pageSize,default=20"`
}

type WithdrawalInfo struct {
	Id          int64  `json:"id"`
	ShopId      int64  `json:"shopId"`
	Amount      int64  `json:"amount"`
	BankInfo    string `json:"bankInfo"`
	Status      int32  `json:"status"`
	AdminId     int64  `json:"adminId"`
	AdminRemark string `json:"adminRemark"`
	CreateTime  int64  `json:"createTime"`
	UpdateTime  int64  `json:"updateTime"`
}

type ListWithdrawalsResp struct {
	Total       int64             `json:"total"`
	Withdrawals []*WithdrawalInfo `json:"withdrawals"`
}

type AdminHandleWithdrawalReq struct {
	Action int32  `json:"action"`
	Remark string `json:"remark,omitempty"`
}

// ===== P2 B-5/B-6: Shop level system =====
type ShopLevelTemplateInfo struct {
	Level          int32   `json:"level"`
	Name           string  `json:"name"`
	MinGmv         int64   `json:"minGmv"`
	MinCreditScore int32   `json:"minCreditScore"`
	MinMonths      int32   `json:"minMonths"`
	MinRating      float64 `json:"minRating"`
	CommissionRate float64 `json:"commissionRate"`
	TrafficBoost   float64 `json:"trafficBoost"`
	Benefits       string  `json:"benefits"`
}

type ListShopLevelsResp struct {
	Levels []*ShopLevelTemplateInfo `json:"levels"`
}

type MyLevelStatusResp struct {
	CurrentLevel          int32                  `json:"currentLevel"`
	CurrentTemplate       *ShopLevelTemplateInfo `json:"currentTemplate"`
	NextTemplate          *ShopLevelTemplateInfo `json:"nextTemplate"`
	CurrentGmv            int64                  `json:"currentGmv"`
	CurrentCreditScore    int32                  `json:"currentCreditScore"`
	CurrentMonths         int32                  `json:"currentMonths"`
	CurrentRating         float64                `json:"currentRating"`
	EligibleForNext       bool                   `json:"eligibleForNext"`
	HasPendingApplication bool                   `json:"hasPendingApplication"`
}

type SubmitLevelApplicationReq struct {
	TargetLevel int32 `json:"targetLevel"`
}

type SubmitLevelApplicationResp struct {
	ApplicationId int64 `json:"applicationId"`
}

type ShopLevelApplicationInfo struct {
	Id           int64  `json:"id"`
	ShopId       int64  `json:"shopId"`
	CurrentLevel int32  `json:"currentLevel"`
	TargetLevel  int32  `json:"targetLevel"`
	Snapshot     string `json:"snapshot"`
	Status       int32  `json:"status"`
	AdminId      int64  `json:"adminId"`
	AdminRemark  string `json:"adminRemark"`
	CreateTime   int64  `json:"createTime"`
}

type ListLevelApplicationsReq struct {
	Status   int32 `form:"status,default=-1"`
	Page     int32 `form:"page,default=1"`
	PageSize int32 `form:"pageSize,default=20"`
}

type ListLevelApplicationsResp struct {
	Total        int64                       `json:"total"`
	Applications []*ShopLevelApplicationInfo `json:"applications"`
}

// ===== P2 G-3: Shop coupons =====
type ShopCouponInfo struct {
	Id              int64  `json:"id"`
	ShopId          int64  `json:"shopId"`
	Code            string `json:"code"`
	Name            string `json:"name"`
	Type            int32  `json:"type"`
	DiscountValue   int64  `json:"discountValue"`
	MinOrderAmount  int64  `json:"minOrderAmount"`
	TotalQuantity   int32  `json:"totalQuantity"`
	ClaimedQuantity int32  `json:"claimedQuantity"`
	PerUserLimit    int32  `json:"perUserLimit"`
	ValidFrom       int64  `json:"validFrom"`
	ValidTo         int64  `json:"validTo"`
	Status          int32  `json:"status"`
	CreateTime      int64  `json:"createTime"`
}

type CreateShopCouponReq struct {
	Code           string `json:"code"`
	Name           string `json:"name"`
	Type           int32  `json:"type"`
	DiscountValue  int64  `json:"discountValue"`
	MinOrderAmount int64  `json:"minOrderAmount,optional"`
	TotalQuantity  int32  `json:"totalQuantity,optional"`
	PerUserLimit   int32  `json:"perUserLimit,optional"`
	ValidFrom      int64  `json:"validFrom"`
	ValidTo        int64  `json:"validTo"`
}

type CreateShopCouponResp struct {
	Id int64 `json:"id"`
}

type ListShopCouponsReq struct {
	Status   int32 `form:"status,default=0"`
	Page     int32 `form:"page,default=1"`
	PageSize int32 `form:"pageSize,default=20"`
}

type ListShopCouponsResp struct {
	Total   int64             `json:"total"`
	Coupons []*ShopCouponInfo `json:"coupons"`
}

type UpdateShopCouponStatusReq struct {
	Status int32 `json:"status"`
}

// ===== P2 G-4: SKU flash discounts =====
type FlashDiscountInfo struct {
	Id            int64 `json:"id"`
	ShopId        int64 `json:"shopId"`
	ProductId     int64 `json:"productId"`
	SkuId         int64 `json:"skuId"`
	OriginalPrice int64 `json:"originalPrice"`
	DiscountPrice int64 `json:"discountPrice"`
	StartTime     int64 `json:"startTime"`
	EndTime       int64 `json:"endTime"`
	Status        int32 `json:"status"`
	CreateTime    int64 `json:"createTime"`
}

type CreateFlashDiscountReq struct {
	ProductId     int64 `json:"productId"`
	SkuId         int64 `json:"skuId"`
	OriginalPrice int64 `json:"originalPrice"`
	DiscountPrice int64 `json:"discountPrice"`
	StartTime     int64 `json:"startTime"`
	EndTime       int64 `json:"endTime"`
}

type CreateFlashDiscountResp struct {
	Id int64 `json:"id"`
}

type ListFlashDiscountsReq struct {
	Status   int32 `form:"status,default=0"`
	Page     int32 `form:"page,default=1"`
	PageSize int32 `form:"pageSize,default=20"`
}

type ListFlashDiscountsResp struct {
	Total     int64                `json:"total"`
	Discounts []*FlashDiscountInfo `json:"discounts"`
}

type ReviewLevelApplicationReq struct {
	Action int32  `json:"action"`
	Remark string `json:"remark,omitempty"`
}

// ===== P2 B-4: shop lifecycle =====
type ShopLifecycleRequestInfo struct {
	Id          int64  `json:"id"`
	ShopId      int64  `json:"shopId"`
	Action      string `json:"action"`
	Reason      string `json:"reason"`
	Status      int32  `json:"status"`
	AdminId     int64  `json:"adminId"`
	AdminRemark string `json:"adminRemark"`
	CreateTime  int64  `json:"createTime"`
}

type SubmitShopLifecycleRequestReq struct {
	Action string `json:"action"`
	Reason string `json:"reason,optional"`
}

type SubmitShopLifecycleRequestResp struct {
	RequestId int64 `json:"requestId"`
}

type ListShopLifecycleRequestsReq struct {
	Status   int32 `form:"status,default=-1"`
	Page     int32 `form:"page,default=1"`
	PageSize int32 `form:"pageSize,default=20"`
}

type ListShopLifecycleRequestsResp struct {
	Total    int64                       `json:"total"`
	Requests []*ShopLifecycleRequestInfo `json:"requests"`
}

type ReviewShopLifecycleRequestReq struct {
	Action int32  `json:"action"`
	Remark string `json:"remark,optional"`
}

// ===== P2 D-3: batch ship orders =====
type BatchShipResult struct {
	Total   int      `json:"total"`
	Success int      `json:"success"`
	Failed  int      `json:"failed"`
	Errors  []string `json:"errors,omitempty"`
}

// ===== P2 C-6: freight templates =====
type FreightTemplateInfo struct {
	Id         int64  `json:"id"`
	ShopId     int64  `json:"shopId"`
	Name       string `json:"name"`
	CalcType   int32  `json:"calcType"`
	FirstValue int32  `json:"firstValue"`
	FirstFee   int64  `json:"firstFee"`
	ExtraValue int32  `json:"extraValue"`
	ExtraFee   int64  `json:"extraFee"`
	Regions    string `json:"regions"`
	IsDefault  bool   `json:"isDefault"`
	Status     int32  `json:"status"`
	CreateTime int64  `json:"createTime"`
}

type CreateFreightTemplateReq struct {
	Name       string `json:"name"`
	CalcType   int32  `json:"calcType"`
	FirstValue int32  `json:"firstValue"`
	FirstFee   int64  `json:"firstFee"`
	ExtraValue int32  `json:"extraValue"`
	ExtraFee   int64  `json:"extraFee"`
	Regions    string `json:"regions,optional"`
	IsDefault  bool   `json:"isDefault,optional"`
}

type CreateFreightTemplateResp struct {
	Id int64 `json:"id"`
}

type ListFreightTemplatesReq struct {
	Page     int32 `form:"page,default=1"`
	PageSize int32 `form:"pageSize,default=20"`
}

type ListFreightTemplatesResp struct {
	Total     int64                  `json:"total"`
	Templates []*FreightTemplateInfo `json:"templates"`
}

type UpdateFreightTemplateReq struct {
	Name      string `json:"name,optional"`
	FirstFee  int64  `json:"firstFee,optional"`
	ExtraFee  int64  `json:"extraFee,optional"`
	IsDefault bool   `json:"isDefault,optional"`
}

// ===== P2 J-4: sensitive words =====
type SensitiveWordInfo struct {
	Id         int64  `json:"id"`
	Word       string `json:"word"`
	Category   string `json:"category"`
	Action     string `json:"action"`
	Status     int32  `json:"status"`
	CreateTime int64  `json:"createTime"`
}

type CreateSensitiveWordReq struct {
	Word     string `json:"word"`
	Category string `json:"category,optional"`
	Action   string `json:"action,optional"`
}

type CreateSensitiveWordResp struct {
	Id int64 `json:"id"`
}

type ListSensitiveWordsReq struct {
	Category string `form:"category,optional"`
	Page     int32  `form:"page,default=1"`
	PageSize int32  `form:"pageSize,default=20"`
}

type ListSensitiveWordsResp struct {
	Total int64                `json:"total"`
	Words []*SensitiveWordInfo `json:"words"`
}

// ===== S2 Refund =====
type RefundItemDTO struct {
	SkuId    int64  `json:"skuId"`
	SkuName  string `json:"skuName"`
	Quantity int32  `json:"quantity"`
	Amount   int64  `json:"amount"`
}

type RefundInfo struct {
	Id                 int64           `json:"id"`
	OrderId            int64           `json:"orderId"`
	OrderNo            string          `json:"orderNo"`
	UserId             int64           `json:"userId"`
	ShopId             int64           `json:"shopId"`
	Amount             int64           `json:"amount"`
	Reason             string          `json:"reason"`
	Evidence           []string        `json:"evidence"`
	Items              []RefundItemDTO `json:"items"`
	Status             int32           `json:"status"`
	MerchantUserId     int64           `json:"merchantUserId"`
	MerchantRemark     string          `json:"merchantRemark"`
	MerchantHandleTime int64           `json:"merchantHandleTime"`
	AdminId            int64           `json:"adminId"`
	AdminRemark        string          `json:"adminRemark"`
	AdminHandleTime    int64           `json:"adminHandleTime"`
	AppealReason       string          `json:"appealReason"`
	AppealTime         int64           `json:"appealTime"`
	RefundNo           string          `json:"refundNo"`
	RefundCompleteTime int64           `json:"refundCompleteTime"`
	CreateTime         int64           `json:"createTime"`
}

type ListRefundsReq struct {
	Status   int32 `form:"status,default=-1"`
	Page     int32 `form:"page,default=1"`
	PageSize int32 `form:"pageSize,default=20"`
}

type ListRefundsResp struct {
	Total    int64         `json:"total"`
	Requests []*RefundInfo `json:"requests"`
}

type ArbitrateRefundReq struct {
	Action int32  `json:"action"` // 1=force_refund 2=final_reject
	Remark string `json:"remark,omitempty"`
}

type MerchantHandleRefundReq struct {
	Action int32  `json:"action"` // 1=approve 2=reject
	Remark string `json:"remark,omitempty"`
}

// ===== S3 account ledger =====

type LedgerEntryDTO struct {
	Id             int64  `json:"id"`
	ShopId         int64  `json:"shopId"`
	Direction      int32  `json:"direction"`
	Category       string `json:"category"`
	Amount         int64  `json:"amount"`
	RunningBalance int64  `json:"runningBalance"`
	OrderId        int64  `json:"orderId"`
	RefundId       int64  `json:"refundId"`
	RefNo          string `json:"refNo"`
	Description    string `json:"description"`
	CreateTime     int64  `json:"createTime"`
}

type ListLedgerReq struct {
	ShopId    int64  `form:"shopId,optional"`
	Category  string `form:"category,optional"`
	StartTime int64  `form:"startTime,optional"`
	EndTime   int64  `form:"endTime,optional"`
	Page      int32  `form:"page,default=1"`
	PageSize  int32  `form:"pageSize,default=20"`
}

type ListLedgerResp struct {
	Entries []LedgerEntryDTO `json:"entries"`
	Total   int64            `json:"total"`
}

type GetLedgerSummaryReq struct {
	ShopId    int64 `form:"shopId,optional"`
	StartTime int64 `form:"startTime,optional"`
	EndTime   int64 `form:"endTime,optional"`
}

type LedgerSummaryDTO struct {
	TotalIncome     int64 `json:"totalIncome"`
	TotalRefund     int64 `json:"totalRefund"`
	TotalCommission int64 `json:"totalCommission"`
	TotalWithdrawal int64 `json:"totalWithdrawal"`
	NetBalance      int64 `json:"netBalance"`
}

type RunReconcileReq struct {
	ShopId int64 `json:"shopId,optional"`
}

type ShopReconcileResultDTO struct {
	ShopId        int64 `json:"shopId"`
	LedgerCredit  int64 `json:"ledgerCredit"`
	LedgerDebit   int64 `json:"ledgerDebit"`
	LedgerNet     int64 `json:"ledgerNet"`
	WalletBalance int64 `json:"walletBalance"`
	WalletFrozen  int64 `json:"walletFrozen"`
	WalletTotal   int64 `json:"walletTotal"`
	Diff          int64 `json:"diff"`
	Passed        bool  `json:"passed"`
}

type ReconcileReportDTO struct {
	TotalChecked int32                    `json:"totalChecked"`
	Passed       int32                    `json:"passed"`
	Failed       int32                    `json:"failed"`
	Results      []ShopReconcileResultDTO `json:"results"`
}
