package logic

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"mall-admin-api/internal/middleware"
	"mall-admin-api/internal/svc"
	"mall-admin-api/internal/types"

	"mall-order-rpc/orderclient"
)

// BatchShipOrders parses a merchant-uploaded CSV (order_id,carrier,tracking_no)
// and ships each row via mall-order-rpc.ShipOrder. Per-row failures are
// collected but do not abort the batch.
func BatchShipOrders(ctx context.Context, svcCtx *svc.ServiceContext, reader io.Reader) (*types.BatchShipResult, error) {
	c, _ := middleware.ClaimsFromContext(ctx)
	if c == nil || c.ShopId <= 0 {
		return nil, errors.New("merchant shop unknown")
	}

	csvReader := csv.NewReader(reader)
	csvReader.FieldsPerRecord = -1
	result := &types.BatchShipResult{Errors: []string{}}

	lineNum := 0
	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		lineNum++
		if err != nil {
			result.Errors = append(result.Errors, fmt.Sprintf("line %d: parse error: %v", lineNum, err))
			result.Failed++
			continue
		}
		if len(row) < 3 {
			result.Errors = append(result.Errors, fmt.Sprintf("line %d: need 3 columns, got %d", lineNum, len(row)))
			result.Failed++
			continue
		}
		// Skip header row.
		if lineNum == 1 && !isDigits(strings.TrimSpace(row[0])) {
			continue
		}
		result.Total++

		orderId, perr := strconv.ParseInt(strings.TrimSpace(row[0]), 10, 64)
		if perr != nil || orderId <= 0 {
			result.Errors = append(result.Errors, fmt.Sprintf("line %d: invalid order_id %q", lineNum, row[0]))
			result.Failed++
			continue
		}
		carrier := strings.TrimSpace(row[1])
		trackingNo := strings.TrimSpace(row[2])
		if _, err := svcCtx.OrderRpc.ShipOrder(ctx, &orderclient.ShipOrderReq{
			Id:         orderId,
			ShopId:     c.ShopId,
			Carrier:    carrier,
			TrackingNo: trackingNo,
		}); err != nil {
			result.Errors = append(result.Errors, fmt.Sprintf("line %d: order %d: %v", lineNum, orderId, err))
			result.Failed++
			continue
		}
		result.Success++
	}
	return result, nil
}

func isDigits(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
