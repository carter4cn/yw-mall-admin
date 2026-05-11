# syntax=docker/dockerfile:1
# Build context: parent dir (go/mall/) so both yw-mall-admin/ and yw-mall/mall-*-rpc/ are reachable.
FROM docker.io/library/golang:1.26-alpine AS builder
WORKDIR /workspace

ARG GOPROXY=
ENV GOPROXY=${GOPROXY:+$GOPROXY}

# Layer 1: go.mod + go.sum — cached when only source changes
COPY yw-mall-admin/go.mod yw-mall-admin/go.sum              ./yw-mall-admin/
COPY yw-mall/mall-common/go.mod yw-mall/mall-common/go.sum  ./yw-mall/mall-common/
COPY yw-mall/mall-user-rpc/go.mod yw-mall/mall-user-rpc/go.sum           ./yw-mall/mall-user-rpc/
COPY yw-mall/mall-shop-rpc/go.mod yw-mall/mall-shop-rpc/go.sum           ./yw-mall/mall-shop-rpc/
COPY yw-mall/mall-product-rpc/go.mod yw-mall/mall-product-rpc/go.sum     ./yw-mall/mall-product-rpc/
COPY yw-mall/mall-order-rpc/go.mod yw-mall/mall-order-rpc/go.sum         ./yw-mall/mall-order-rpc/
COPY yw-mall/mall-logistics-rpc/go.mod yw-mall/mall-logistics-rpc/go.sum ./yw-mall/mall-logistics-rpc/
COPY yw-mall/mall-workflow-rpc/go.mod yw-mall/mall-workflow-rpc/go.sum   ./yw-mall/mall-workflow-rpc/
COPY yw-mall/mall-rule-rpc/go.mod yw-mall/mall-rule-rpc/go.sum           ./yw-mall/mall-rule-rpc/
COPY yw-mall/mall-review-rpc/go.mod yw-mall/mall-review-rpc/go.sum       ./yw-mall/mall-review-rpc/
COPY yw-mall/mall-risk-rpc/go.mod yw-mall/mall-risk-rpc/go.sum           ./yw-mall/mall-risk-rpc/
COPY yw-mall/mall-payment-rpc/go.mod yw-mall/mall-payment-rpc/go.sum     ./yw-mall/mall-payment-rpc/
COPY yw-mall/mall-activity-rpc/go.mod yw-mall/mall-activity-rpc/go.sum   ./yw-mall/mall-activity-rpc/

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    cd /workspace/yw-mall-admin && go mod download

# Layer 2: full source
COPY yw-mall-admin                  ./yw-mall-admin
COPY yw-mall/mall-common            ./yw-mall/mall-common
COPY yw-mall/mall-user-rpc          ./yw-mall/mall-user-rpc
COPY yw-mall/mall-shop-rpc          ./yw-mall/mall-shop-rpc
COPY yw-mall/mall-product-rpc       ./yw-mall/mall-product-rpc
COPY yw-mall/mall-order-rpc         ./yw-mall/mall-order-rpc
COPY yw-mall/mall-logistics-rpc     ./yw-mall/mall-logistics-rpc
COPY yw-mall/mall-workflow-rpc      ./yw-mall/mall-workflow-rpc
COPY yw-mall/mall-rule-rpc          ./yw-mall/mall-rule-rpc
COPY yw-mall/mall-review-rpc        ./yw-mall/mall-review-rpc
COPY yw-mall/mall-risk-rpc          ./yw-mall/mall-risk-rpc
COPY yw-mall/mall-payment-rpc       ./yw-mall/mall-payment-rpc
COPY yw-mall/mall-activity-rpc      ./yw-mall/mall-activity-rpc

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    mkdir -p /out && \
    cd /workspace/yw-mall-admin && \
    CGO_ENABLED=0 GOOS=linux go build -trimpath -o /out/server . && \
    cp -r etc /out/etc

FROM docker.io/library/alpine:3.21
RUN apk add --no-cache ca-certificates tzdata
ENV TZ=Asia/Shanghai
WORKDIR /app
COPY --from=builder /out/server           ./server
COPY --from=builder /out/etc              ./etc
COPY yw-mall/docker-entrypoint.sh         ./entrypoint.sh
RUN chmod +x ./entrypoint.sh
ENTRYPOINT ["./entrypoint.sh"]
