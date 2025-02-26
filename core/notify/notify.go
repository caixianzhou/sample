// Copyright 2021 Inc. All rights reserved.

// Package notify  API Go SDK 商户通知处理库
package notify

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/core/auth/validators"
)

// Handler 通知 Handler
type Handler struct {
	Validator validators.MidasBuyNotifyValidator
}

// ParseNotifyRequest 从 HTTP 请求(http.Request) 中解析 通知(notify.Request)
func (h *Handler) ParseNotifyRequest(ctx context.Context, request *http.Request) (
	*Notification, error,
) {
	if err := h.Validator.Validate(ctx, request); err != nil {
		return nil, fmt.Errorf("not valid midasbuy notify request: %v", err)
	}

	body, err := getRequestBody(request)
	if err != nil {
		return nil, err
	}

	ret := &Notification{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, fmt.Errorf("parse request body error: %v", err)
	}
	return ret, nil
}

// ParseNotifyRequestWithResource 从 HTTP 请求(http.Request) 中解析到resource
func (h *Handler) ParseNotifyRequestWithResource(ctx context.Context, request *http.Request, resource interface{}) error {
	if err := h.Validator.Validate(ctx, request); err != nil {
		return fmt.Errorf("not valid midasbuy notify request: %v", err)
	}

	body, err := getRequestBody(request)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(body, resource); err != nil {
		return fmt.Errorf("parse request body error: %v", err)
	}
	return nil
}

func getRequestBody(request *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, fmt.Errorf("read request body err: %v", err)
	}

	_ = request.Body.Close()
	request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return body, nil
}
