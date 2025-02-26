// Copyright 2024 Tencent Inc. All rights reserved.

// Package credentials Go SDK 请求报文头 Authorization 信息生成器
package credentials

import (
	"context"
	"fmt"
	"time"

	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/core/auth"
	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/core/consts"
	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/utils"
)

// MidasBuyCredentials MidasBuy请求报文头 Authorization 信息生成器
type MidasBuyCredentials struct {
	Signer auth.Signer // 数字签名生成器
}

// GenerateAuthorizationHeader 生成请求报文头中的 Authorization 信息，详见：
func (c *MidasBuyCredentials) GenerateAuthorizationHeader(
	ctx context.Context, authIDType consts.AuthIDType, method, canonicalURL, signBody string,
) (string, error) {
	if c.Signer == nil {
		return "", fmt.Errorf("you must init MidasBuyCredentials with signer")
	}
	nonce, err := utils.GenerateNonce()
	if err != nil {
		return "", err
	}
	timestamp := time.Now().Unix()
	message := fmt.Sprintf(consts.SignatureMessageFormat, method, canonicalURL, timestamp, nonce, signBody)
	signatureResult, err := c.Signer.Sign(ctx, message)
	if err != nil {
		return "", err
	}
	authorization := fmt.Sprintf(
		consts.HeaderAuthorizationFormat, c.getAuthorizationType(),
		signatureResult.MchID, authIDType, nonce, timestamp, signatureResult.CertificateSerialNo,
		signatureResult.Signature,
	)
	return authorization, nil
}

func (c *MidasBuyCredentials) getAuthorizationType() string {
	return consts.HeaderAuthorizationPrefix + c.Signer.Algorithm()
}
