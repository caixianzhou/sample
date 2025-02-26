// Copyright 2024 Tencent Inc. All rights reserved.

package validators

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/core/consts"
	"github.com/stretchr/testify/assert"
)

type mockVerifier struct {
}

func (v *mockVerifier) GetSerial(ctx context.Context) (serial string, err error) {
	return "SERIAL1234567890", nil
}

func (v *mockVerifier) Verify(ctx context.Context, serialNumber string, message string, signature string) error {
	if "["+serialNumber+"-"+message+"]" == signature {
		return nil
	}

	return fmt.Errorf("verification failed")
}

func TestNullValidator_Validate(t *testing.T) {
	nullValidator := NullValidator{}

	assert.NoError(t, nullValidator.Validate(context.Background(), &http.Response{}))
	assert.NoError(t, nullValidator.Validate(context.Background(), nil))
}

func TestMidasBuyNotifyValidator_Validate(t *testing.T) {
	mockTimestamp := time.Now().Unix()
	mockTimestampStr := fmt.Sprintf("%d", mockTimestamp)

	validator := NewMidasBuyNotifyValidator(&mockVerifier{})

	request := httptest.NewRequest("Post", "http://127.0.0.1", ioutil.NopCloser(bytes.NewBuffer([]byte("BODY"))))
	request.Header = http.Header{
		consts.MidasBuySignature: {
			"[SERIAL1234567890-" + mockTimestampStr + "\nNONCE1234567890\nBODY\n]",
		},
		consts.MidasBuySerial:    {"SERIAL1234567890"},
		consts.MidasBuyTimestamp: {mockTimestampStr},
		consts.MidasBuyNonce:     {"NONCE1234567890"},
		consts.RequestID:         {"any-request-id"},
	}

	err := validator.Validate(context.Background(), request)
	assert.NoError(t, err)
}
