// Copyright 2024 Tencent Inc. All rights reserved.

package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// APIError MidasBuy API 标准错误结构
type APIError struct {
	StatusCode int         // 应答报文的 HTTP 状态码
	Header     http.Header // 应答报文的 Header 信息
	Body       string      // 应答报文的 Body 原文
	Name       string      `json:"name"`             // 应答报文的 Body 解析后的错误码信息，仅不符合预期/发生系统错误时存在
	Message    string      `json:"message"`          // 应答报文的 Body 解析后的文字说明信息，仅不符合预期/发生系统错误时存在
	Detail     interface{} `json:"detail,omitempty"` // 应答报文的 Body 解析后的详细信息，仅不符合预期/发生系统错误时存在
	DebugID    string      `json:"debug_id"`         // 应答报文的 Body 解析后的调试信息，仅不符合预期/发生系统错误时存在
}

// Error 输出 APIError
func (e *APIError) Error() string {
	var buf bytes.Buffer
	_, _ = fmt.Fprintf(&buf, "error http response:[StatusCode: %d Name: \"%s\"", e.StatusCode, e.Name)
	if e.Message != "" {
		_, _ = fmt.Fprintf(&buf, "\nMessage: %s", e.Message)
	}
	if e.Detail != nil {
		var detailBuf bytes.Buffer
		enc := json.NewEncoder(&detailBuf)
		enc.SetIndent("", "  ")
		if err := enc.Encode(e.Detail); err == nil {
			_, _ = fmt.Fprint(&buf, "\nDetail:")
			_, _ = fmt.Fprintf(&buf, "\n%s", strings.TrimSpace(detailBuf.String()))
		}
	}
	if len(e.Header) > 0 {
		_, _ = fmt.Fprint(&buf, "\nHeader:")
		for key, value := range e.Header {
			_, _ = fmt.Fprintf(&buf, "\n - %v=%v", key, value)
		}
	}
	if e.DebugID != "" {
		_, _ = fmt.Fprintf(&buf, "\nDebugID: %s", e.DebugID)
	}
	_, _ = fmt.Fprintf(&buf, "]")
	return buf.String()
}

// IsAPIError 判断当前 error 是否为特定 Code 的 *APIError
//
// 类型为其他 error 或 Code 不匹配时均返回 false
func IsAPIError(err error, name string) bool {
	var ne *APIError
	if errors.As(err, &ne) {
		return ne.Name == name
	}
	return false
}
