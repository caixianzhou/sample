// Copyright 2024 Tencent Inc. All rights reserved.

package core

import (
	"fmt"
	"net/http"

	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/core/auth"
)

// DialSettings MidasBuy API Go SDK core.Client 需要的配置信息
type DialSettings struct {
	HTTPClient *http.Client   // 自定义所使用的 HTTPClient 实例
	Signer     auth.Signer    // 签名器
	Validator  auth.Validator // 应答包签名校验器
}

// Validate 校验请求配置是否有效
func (ds *DialSettings) Validate() error {
	if ds.Validator == nil {
		return fmt.Errorf("validator is required for Client")
	}
	if ds.Signer == nil {
		return fmt.Errorf("signer is required for Client")
	}
	return nil
}
