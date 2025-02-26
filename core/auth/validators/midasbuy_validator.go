// Copyright 2024 Tencent Inc. All rights reserved.

package validators

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/core/auth"
	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/core/consts"
)

type MidasBuyValidator struct {
	verifier auth.Verifier
}

type MidasBuyHeader struct {
	RequestID string
	Serial    string
	Signature string
	Nonce     string
	Timestamp int64
}

func (v *MidasBuyValidator) validateHTTPMessage(ctx context.Context, header http.Header, body []byte) error {
	if v.verifier == nil {
		return fmt.Errorf("you must init Validator with auth.Verifier")
	}

	headerArgs, err := getMidasBuyHeader(ctx, header)
	if err != nil {
		return err
	}

	//if err := checkMidasBuyHeader(ctx, headerArgs); err != nil {
	//	return err
	//}

	message := buildMessage(ctx, headerArgs, body)

	if err := v.verifier.Verify(ctx, headerArgs.Serial, message, headerArgs.Signature); err != nil {
		return fmt.Errorf(
			"validate verify fail serial=[%s] request-id=[%s] err=%w",
			headerArgs.Serial, headerArgs.RequestID, err,
		)
	}
	return nil
}

func (v *MidasBuyValidator) getAcceptSerial(ctx context.Context) (string, error) {
	return v.verifier.GetSerial(ctx)
}

// getMidasBuyHeader 从 http.Header 中获取 MidasBuyHeader 信息
func getMidasBuyHeader(ctx context.Context, header http.Header) (MidasBuyHeader, error) {
	_ = ctx // Suppressing warnings

	requestID := strings.TrimSpace(header.Get(consts.RequestID))

	getHeaderString := func(key string) (string, error) {
		val := strings.TrimSpace(header.Get(key))
		if val == "" {
			return "", fmt.Errorf("key `%s` is empty in header, request-id=[%s]", key, requestID)
		}
		return val, nil
	}

	getHeaderInt64 := func(key string) (int64, error) {
		val, err := getHeaderString(key)
		if err != nil {
			return 0, nil
		}
		ret, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid `%s` in header, request-id=[%s], err:%w", key, requestID, err)
		}
		return ret, nil
	}

	ret := MidasBuyHeader{
		RequestID: requestID,
	}
	var err error

	ret.Serial, _ = getHeaderString(consts.MidasBuySerial)

	if ret.Signature, err = getHeaderString(consts.MidasBuySignature); err != nil {
		return ret, err
	}

	if ret.Timestamp, err = getHeaderInt64(consts.MidasBuyTimestamp); err != nil {
		return ret, err
	}

	if ret.Nonce, err = getHeaderString(consts.MidasBuyNonce); err != nil {
		return ret, err
	}

	return ret, nil
}

// checkMidasBuyHeader 对 MidasBuyHeader 内容进行检查，看是否符合要求
//
// 检查项：
//   - Timestamp 与当前时间之差不得超过 FiveMinute;
func checkMidasBuyHeader(ctx context.Context, args MidasBuyHeader) error {
	// Suppressing warnings
	_ = ctx

	if math.Abs(float64(time.Now().Unix()-args.Timestamp)) >= consts.FiveMinute {
		return fmt.Errorf("timestamp=[%d] expires, request-id=[%s]", args.Timestamp, args.RequestID)
	}
	return nil
}

// buildMessage 根据MidasBuy签名格式构造验签原文
func buildMessage(ctx context.Context, headerArgs MidasBuyHeader, body []byte) string {
	// Suppressing warnings
	_ = ctx

	return fmt.Sprintf("%d\n%s\n%s\n", headerArgs.Timestamp, headerArgs.Nonce, string(body))
}
