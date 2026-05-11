# yw-mall-admin

yw-mall 后台管理 HTTP API 网关 — 服务管理员与店家两类后台用户。

> C 端商城走 `mall-api` (`/api/v1/...`)，本服务专门承接后台流量：`/admin/v1/...`（平台管理员）与 `/merchant/v1/...`（店家）。

---

## 架构定位

```
┌─────────────────┐         ┌──────────────────┐         ┌──────────────────┐
│ Admin Portal FE │ ──────▶ │ yw-mall-admin    │ ──gRPC─▶│ mall-*-rpc       │
│ Merchant Portal │  HTTP   │ (this service)   │  (zrpc) │ (12 services)    │
└─────────────────┘  :18999 └──────────────────┘         └──────────────────┘
                                    │                              │
                                    ▼                              ▼
                            ┌───────────────┐              ┌──────────────┐
                            │  etcd 配置中心  │              │ ProxySQL/    │
                            │  /config/dev/…  │              │ MySQL/Redis  │
                            └───────────────┘              └──────────────┘
```

- 路由分组：`/admin/v1/login` 公开，其余 admin 路由经 `AdminAuth` + `OpLog`；merchant 同理。
- 鉴权：JWT 中携带 `Uid / Role / ShopId / Perms`；同一 secret，按 path 前缀分流。
- 配置：启动从 etcd `/config/${APP_ENV}/yw-mall/admin-api` 加载（fallback 到 `etc/admin.yaml`）。

---

## 后端依赖

下游 RPC 全部通过 etcd 服务发现：

| Service | 用途 |
|---|---|
| `mall-user-rpc` | admin 账号、用户列表/封禁 |
| `mall-shop-rpc` | 入驻审核、店家 CRUD、信用分、等级、生命周期 |
| `mall-product-rpc` | 商品审核 + 商家 CRUD + 库存 |
| `mall-order-rpc` | 店家订单、发货、退款 |
| `mall-review-rpc` | 评论列表、回复、删除申请 |
| `mall-risk-rpc` | 投诉工单、限制、敏感词 |
| `mall-activity-rpc` | 活动、优惠券、限时折扣 |
| `mall-rule-rpc` | 低代码活动规则 |
| `mall-payment-rpc` | 钱包、提现、账单 |
| `mall-logistics-rpc` | 运费模板 |
| `mall-workflow-rpc` | 审批流（入驻、提现） |

---

## 路由总览

### `/admin/v1/...`（平台管理员）

| 模块 | Endpoints |
|---|---|
| 登录 | `POST /login` |
| 账号 | `POST /accounts`, `GET /accounts` |
| 入驻审核 | `GET /shop-applications`, `GET /shop-applications/:id`, `POST /shop-applications/:id/review` |
| 店家管理 | `GET /shops`, `POST /shops/:id/status`, `POST /shops/:id/credit` (自动触发信用阈值规则) |
| 商品审核 | `GET /products/review`, `POST /products/:id/review` |
| 用户管理 | `GET /users`, `POST /users/:id/status` |
| 评论 | `GET /reviews/delete-requests`, `POST /reviews/delete-requests/:id/handle` |
| 投诉 | `GET /complaints`, `GET /complaints/:id`, `POST /complaints/:id/handle` |
| 限制 | `POST /shops/:id/restrictions`, `GET /shops/:id/restrictions`, `DELETE /shops/:id/restrictions/:rid` |
| 活动 | `GET/POST /activities`, `PUT /activities/:id`, `POST /activities/:id/status` |
| 规则 | `GET/POST /rules`, `POST /rules/validate`, `POST /activity-rules` |
| 提现 | `GET /withdrawals`, `POST /withdrawals/:id/handle` |
| 等级 | `GET /level-applications`, `POST /level-applications/:id/review` |
| 生命周期 | `GET /shop-lifecycle-requests`, `POST /shop-lifecycle-requests/:id/review` |
| 敏感词 | `POST/GET /sensitive-words`, `DELETE /sensitive-words/:id` |

### `/merchant/v1/...`（店家）

| 模块 | Endpoints |
|---|---|
| 登录 / 入驻 | `POST /login`, `POST /apply`, `GET /apply/:id` |
| 店铺 | `GET/PUT /shop`, `POST /shop/lifecycle` |
| 商品 | `GET/POST /products`, `GET/PUT /products/:id`, `POST /products/:id/status`, `POST /products/:id/stock` |
| 订单 | `GET /orders`, `GET /orders/:id`, `POST /orders/:id/ship`, `POST /orders/:id/reject-refund`, `POST /orders/batch-ship`（CSV） |
| 评论 | `GET /reviews`, `POST /reviews/:id/delete-request` |
| 投诉 | `POST /complaints` |
| 活动 | `GET /activities` |
| 钱包 | `GET /wallet`, `GET /wallet/bills`, `POST /wallet/withdraw`, `GET /wallet/withdrawals` |
| 等级 | `GET /shop-levels`, `GET /shop/level-status`, `POST /shop/level/apply` |
| 优惠券 | `POST/GET /coupons`, `POST /coupons/:id/status` |
| 限时折扣 | `POST/GET /flash-discounts`, `POST /flash-discounts/:id/cancel` |
| 运费模板 | `POST/GET /freight-templates`, `GET/PUT/DELETE /freight-templates/:id` |

---

## 启动

### 容器（推荐 — 配合 yw-mall-deploy）

```bash
# 在 yw-mall-deploy/ 目录下
podman-compose build mall-admin-api
podman-compose up -d mall-admin-api
```

服务监听 `:18999`，由 `docker-entrypoint.sh` 把 `127.0.0.1` 重写为 compose 服务名。

### 本地（开发）

```bash
go build -o /tmp/admin-api .
./tmp/admin-api -f etc/admin.yaml
```

需先把 etcd / proxysql / redis 用 compose 启起来。

---

## 配置

`etc/admin.yaml`：
```yaml
Name: mall-admin-api
Host: 0.0.0.0
Port: 18999

JwtSecret: <change-me>            # 与 mall-api 完全独立

UserRpc:   { Etcd: { Hosts: [127.0.0.1:2379], Key: yw-mall/user-rpc } }
ShopRpc:   { Etcd: { Hosts: [127.0.0.1:2379], Key: yw-mall/shop-rpc } }
# … 其余下游 RPC 同样格式
```

部署时通过 `scripts/config-push.sh` 把 yaml 推到 etcd：`/config/${APP_ENV}/yw-mall/admin-api`。

---

## 默认账号

容器 + db-init 启动后会播种 super admin：

```
username: admin
password: admin123
```

> 生产环境请通过 `POST /admin/v1/accounts` 新建账号后立即停用此默认账号。

---

## 鉴权

### 登录返回

```json
{ "token": "eyJ...", "uid": 1, "role": "super_admin", "shopId": 0 }
```

### 请求

```http
GET /admin/v1/level-applications?status=-1&page=1&page_size=10
Authorization: Bearer eyJ...
```

### Claims

```go
type Claims struct {
    Uid    int64
    Role   string   // "admin" / "merchant" / "super_admin"
    ShopId int64    // merchant 才有
    Perms  []string // RBAC
}
```

---

## 中间件

| 中间件 | 范围 | 行为 |
|---|---|---|
| `AdminAuth` | `/admin/v1/*`（除 login） | 校验 JWT，注入 Claims；role != admin/super_admin → 401 |
| `MerchantAuth` | `/merchant/v1/*`（除 login/apply） | 校验 JWT；role != merchant → 401；注入 ShopId |
| `OpLog` | 所有 protected 路由 | 结构化日志写操作（actor / method / path / status / body 摘要） |
| `Rbac` | 待启用 | 按 Perms 字段二级授权 |

---

## 开发

### 增加一个新端点

1. `internal/types/types.go` — 加请求/响应 struct
2. `internal/handler/routes.go` — 加 route 入口
3. `internal/handler/handlers.go` — 加 handler 函数（参考已有的 `listLevelApplicationsHandler` 等）
4. `internal/logic/<area>_logic.go` — 业务逻辑（薄壳：解析参数 → 调 RPC → 拼响应）
5. 如需新 RPC 依赖：`config.go` + `svc/servicecontext.go` + `etc/admin.yaml`，并执行 `scripts/config-push.sh` 推 etcd

### 路径参数

务必用 `parseId(r)`（基于 `pathvar.Vars(r)`）取 `:id`/`:rid`，**不要**用 `httpx.Parse` — 后者会消费 body，与下游再次 Parse 冲突。

### proto 变更

本服务自身不持有 proto；都在下游 RPC 服务的 `mall-common/proto/`。修改后：

```bash
cd /home/carter/workspace/go/mall/yw-mall/mall-{name}-rpc
protoc --go_out=. --go-grpc_out=. \
       --proto_path=. --proto_path=../mall-common/proto \
       ../mall-common/proto/{name}/{name}.proto
# 然后手动更新 {name}service/{name}service.go 暴露新方法
# 最后回本仓库 go mod tidy
```

---

## 关联仓库

| Repo | 内容 |
|---|---|
| [yw-mall](https://github.com/carter4cn/yw-mall) | 12 个 mall-*-rpc 后端服务 + mall-api C 端网关 |
| [yw-mall-deploy](https://github.com/carter4cn/yw-mall-deploy) | podman-compose 编排 + 启动脚本 + DDL bootstrap |
| yw-mall-admin-fe | 后台前端（管理员 + 店家门户） |
| mall-frontend | C 端商城前端 |

---

## PRD

完整需求文档：[docs/prd_admin_merchant_portal.md](docs/prd_admin_merchant_portal.md)

PRD 涵盖 Epic A 身份权限 → J 运营工具共 10 个 Epic、47 个 Story，按 P0/P1/P2 优先级落地。
