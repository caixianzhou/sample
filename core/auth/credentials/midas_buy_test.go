// Copyright 2024 Tencent Inc. All rights reserved.

package credentials

import (
	"context"
	"fmt"
	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/core/consts"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/core/auth"
	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/utils"
	"github.com/agiledragon/gomonkey"
	"github.com/stretchr/testify/require"
)

type mockSigner struct {
	MchID               string
	CertificateSerialNo string
}

func (s *mockSigner) Sign(_ context.Context, message string) (*auth.SignatureResult, error) {
	result := &auth.SignatureResult{
		MchID:               s.MchID,
		CertificateSerialNo: s.CertificateSerialNo,
		Signature:           "Sign:" + message,
	}
	return result, nil
}

func (s *mockSigner) Algorithm() string {
	return "Mock"
}

type mockErrorSigner struct {
}

func (s *mockErrorSigner) Sign(_ context.Context, message string) (*auth.SignatureResult, error) {
	return nil, fmt.Errorf("mock sign error")
}

func (s *mockErrorSigner) Algorithm() string {
	return "ErrorMock"
}

const (
	testMchID             = "1234567890"
	testCertificateSerial = "0123456789ABC"
	mockNonce             = "A1B2C3D4E5F6G7"
	mockTimestamp         = 1624523846
)

func TestMidasBuyCredentials_GenerateAuthorizationHeader(t *testing.T) {
	patches := gomonkey.NewPatches()
	defer patches.Reset()

	patches.ApplyFunc(
		utils.GenerateNonce, func() (string, error) {
			return mockNonce, nil
		},
	)
	patches.ApplyFunc(
		time.Now, func() time.Time {
			return time.Unix(mockTimestamp, 0)
		},
	)

	signer := mockSigner{
		MchID:               testMchID,
		CertificateSerialNo: testCertificateSerial,
	}

	type args struct {
		signer auth.Signer

		ctx          context.Context
		authIDType   consts.AuthIDType
		method       string
		canonicalURL string
		signBody     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    string
	}{
		{
			name: "gen success without body",
			args: args{
				signer: &signer,

				ctx:          context.Background(),
				authIDType:   consts.AuthIDTypeAppID,
				method:       "GET",
				canonicalURL: "/v3/certificates",
				signBody:     "",
			},
			wantErr: false,
			want: `TXGW-Mock auth_id="1234567890",auth_id_type="APP_ID",nonce_str="A1B2C3D4E5F6G7",timestamp="1624523846",` +
				`serial_no="0123456789ABC",signature=` +
				"\"Sign:GET\n/v3/certificates\n1624523846\nA1B2C3D4E5F6G7\n\n\"",
		},
		{
			name: "gen success with body",
			args: args{
				signer: &signer,

				ctx:          context.Background(),
				authIDType:   consts.AuthIDTypeAppID,
				method:       "POST",
				canonicalURL: "/v3/certificates",
				signBody:     "Hello World!\n",
			},
			wantErr: false,
			want: `TXGW-Mock auth_id="1234567890",auth_id_type="APP_ID",nonce_str="A1B2C3D4E5F6G7",timestamp="1624523846",` +
				`serial_no="0123456789ABC",signature=` +
				"\"Sign:POST\n/v3/certificates\n1624523846\nA1B2C3D4E5F6G7\nHello World!\n\n\"",
		},
		{
			name: "gen error wihout signer",
			args: args{
				signer: nil,

				ctx:          context.Background(),
				authIDType:   consts.AuthIDTypeAppID,
				method:       "post",
				canonicalURL: "/v3/certificates",
				signBody:     "Hello World!\n",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				credential := MidasBuyCredentials{Signer: tt.args.signer}

				authorization, err := credential.GenerateAuthorizationHeader(
					tt.args.ctx, tt.args.authIDType, tt.args.method, tt.args.canonicalURL, tt.args.signBody,
				)
				require.Equal(t, tt.wantErr, err != nil)
				require.Equal(t, tt.want, authorization)
			},
		)
	}
}

func TestMidasBuyCredentials_GenerateAuthorizationHeaderErrorGenerateNonce(t *testing.T) {
	patches := gomonkey.NewPatches()
	defer patches.Reset()

	mockGenerateNonceErr := fmt.Errorf("generate nonce error")

	patches.ApplyFunc(
		utils.GenerateNonce, func() (string, error) {
			return "", mockGenerateNonceErr
		},
	)

	signer := mockSigner{
		MchID:               testMchID,
		CertificateSerialNo: testCertificateSerial,
	}
	credential := MidasBuyCredentials{Signer: &signer}

	authorization, err := credential.GenerateAuthorizationHeader(context.Background(), consts.AuthIDTypeAppID, "GET", "/v3/certificates", "")
	require.Error(t, err)
	assert.Empty(t, authorization)
}

func TestMidasBuyCredentials_GenerateAuthorizationHeaderErrorSigner(t *testing.T) {
	patches := gomonkey.NewPatches()
	defer patches.Reset()

	patches.ApplyFunc(
		utils.GenerateNonce, func() (string, error) {
			return mockNonce, nil
		},
	)
	patches.ApplyFunc(
		time.Now, func() time.Time {
			return time.Unix(mockTimestamp, 0)
		},
	)

	signer := mockErrorSigner{}
	credential := MidasBuyCredentials{Signer: &signer}
	authorization, err := credential.GenerateAuthorizationHeader(context.Background(), consts.AuthIDTypeAppID, "GET", "/v3/certificates", "")
	require.Error(t, err)
	assert.Empty(t, authorization)
}
