// Copyright 2024 Tencent Inc. All rights reserved.

package validators

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/core/auth"
)

// MidasBuyNotifyValidator MidasBuy API 通知请求报文验证器
type MidasBuyNotifyValidator struct {
	MidasBuyValidator
}

// Validate 对接收到的MidasBuy API 通知请求报文进行验证
func (v *MidasBuyNotifyValidator) Validate(ctx context.Context, request *http.Request) error {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return fmt.Errorf("read request body err: %v", err)
	}

	_ = request.Body.Close()
	request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return v.validateHTTPMessage(ctx, request.Header, body)
}

// NewMidasBuyNotifyValidator 使用 auth.Verifier 初始化一个 MidasBuyNotifyValidator
func NewMidasBuyNotifyValidator(verifier auth.Verifier) *MidasBuyNotifyValidator {
	return &MidasBuyNotifyValidator{
		MidasBuyValidator{verifier: verifier},
	}
}
