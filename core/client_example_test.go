// Copyright 2024 Tencent Inc. All rights reserved.

package core_test

import (
	"context"
	"encoding/json"
	"log"
	"mime/multipart"
	"net/http"

	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/core"
	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/core/auth"
	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/core/consts"
	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/core/option"
)

func ExampleNewClient_fully_customized() {
	var (
		signer           auth.Signer  // 自定义实现 auth.Signer 接口的实例
		customHTTPClient *http.Client // 自定义 HTTPClient
	)

	client, err := core.NewClient(
		context.Background(),
		// 使用自定义 Signer 初始化 MidasBuy签名器
		option.WithSigner(signer),
		// 使用自定义 HTTPClient
		option.WithHTTPClient(customHTTPClient),
	)
	if err != nil {
		log.Printf("new MidasBuy client err:%s", err.Error())
		return
	}
	// 接下来使用 client 进行请求发送
	_ = client
}

func ExampleCreateFormField() {
	var w multipart.Writer

	meta := map[string]string{
		"filename": "sample.jpg",
		"sha256":   "5944758444f0af3bc843e39b611a6b0c8c38cca44af653cd461b5765b71dc3f8",
	}

	metaBytes, err := json.Marshal(meta)
	if err != nil {
		// TODO: 处理错误
		return
	}

	err = core.CreateFormField(&w, "meta", consts.ApplicationJSON, metaBytes)
	if err != nil {
		// TODO: 处理错误
	}
}

func ExampleCreateFormFile() {
	var w multipart.Writer

	var fileContent []byte

	err := core.CreateFormFile(&w, "sample.jpg", consts.ImageJPG, fileContent)
	if err != nil {
		// TODO: 处理错误
	}
}
