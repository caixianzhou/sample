// Copyright 2024 Tencent Inc. All rights reserved.

// Package services MidasBuy API Go SDK 服务列表
package services

import (
	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/core"
	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/core/consts"
)

// Service MidasBuy API Go SDK 服务类型
type Service struct {
	Client *core.Client
	Env    consts.Env
}
