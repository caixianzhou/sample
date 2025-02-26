// Copyright 2024 Tencent Inc. All rights reserved.

// Package auth MidasBuy API Go SDK 安全验证相关接口
package auth

import (
	"context"
	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/core/consts"
)

// Credential 请求报文头 Authorization 信息生成器
type Credential interface {
	GenerateAuthorizationHeader(
		ctx context.Context, authIDType consts.AuthIDType, method, canonicalURL, signBody string,
	) (string, error)
}
