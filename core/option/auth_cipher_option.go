// Copyright 2024 Tencent Inc. All rights reserved.

package option

import (
	"crypto/rsa"
	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/core"
	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/core/auth/signers"
	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/core/auth/validators"
)

type withAuthCipherOption struct{ settings core.DialSettings }

// Apply 设置 core.DialSettings 的 Signer、Validator 以及 Cipher
func (w withAuthCipherOption) Apply(o *core.DialSettings) error {
	o.Signer = w.settings.Signer
	o.Validator = w.settings.Validator
	return nil
}

// WithMidasBuyPrivateKeyAuth 一键初始化 Client，使其具备「签名」能力。
func WithMidasBuyPrivateKeyAuth(
	mchID, certificateSerialNo string, privateKey *rsa.PrivateKey,
) core.ClientOption {
	return withAuthCipherOption{
		settings: core.DialSettings{
			Signer: &signers.SHA256WithRSASigner{
				MchID:               mchID,
				CertificateSerialNo: certificateSerialNo,
				PrivateKey:          privateKey,
			},
			Validator: &validators.NullValidator{},
		},
	}
}
