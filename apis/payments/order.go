package payments

import (
	"context"
	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/core"
	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/core/consts"
	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/services"
	nethttp "net/http"
	neturl "net/url"
)

type OrderService services.Service

// QueryPaymentOrder 订单查询
// 商户在接收回调后，可通过订单查询接口，获取支付结果
func (s *OrderService) QueryPaymentOrder(
	ctx context.Context, req QueryPaymentOrderInfoRequest,
) (resp *QueryPaymentOrderInfoResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := string(s.Env) + consts.OrderQueryPath
	// Make sure All Required Params are properly set

	// Setup Body Params
	localVarPostBody = req

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = s.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams,
		localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract QueryPaymentOrderInfoResponse from Http Response
	resp = &QueryPaymentOrderInfoResponse{}
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
