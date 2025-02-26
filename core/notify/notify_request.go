// Copyright 2021 Inc. All rights reserved.

package notify

import "encoding/json"

// Notification The notification event.
type Notification struct {
	// The ID of the webhook.
	Id string `json:"id,omitempty"`
	// The date and time when the webhook event notification was created, in Internet date and time format.
	// see: https://tools.ietf.org/html/rfc3339#section-5.6
	CreateTime string `json:"create_time,omitempty"`
	// The time
	UpdateTime string `json:"update_time,omitempty"`
	// The notify resource.
	Resource json.RawMessage `json:"resource,omitempty"`
	// A summary description for the event notification.
	Summary string `json:"summary,omitempty"`
	// The name of the resource related to the webhook notification event.
	ResourceType ResourceType `json:"resource_type,omitempty"`
	// The resource version in the webhook notification.
	ResourceVersion string `json:"resource_version,omitempty"`
	// The event version in the webhook notification.
	EventVersion string `json:"event_version,omitempty"`
	// The event that triggered the webhook event notification.
	EventType EventType `json:"event_type,omitempty"`
}

type EventType string

const (
	PaymentOrderStatusUpdate EventType = "PAYMENT_ORDER_STATUS_UPDATE" // 订单状态更新通知
	PromotionOrderFinished   EventType = "PROMOTION_ORDER_FINISHED"    // 营销完成通知
	ServerValidate           EventType = "SERVER_VALIDATE"             // 查询区服信息
	UserValidate             EventType = "USER_VALIDATE"               // 查询角色信息
	ProductValidate          EventType = "PRODUCT_VALIDATE"            // 查询角色交易资格
)

type ResourceType string

const (
	ResourceTypeOrder     ResourceType = "RESOURCE_TYPE_ORDER"     // 交易订单资源
	ResourceTypePromotion ResourceType = "RESOURCE_TYPE_PROMOTION" // 营销资源
	ResourceTypeUser      ResourceType = "RESOURCE_TYPE_USER"      // 角色资源
	ResourceTypeServer    ResourceType = "RESOURCE_TYPE_SERVER"    // 区服资源
	ResourceTypeProduct   ResourceType = "RESOURCE_TYPE_PRODUCT"   // 资格
)

type Response struct {
	// 商户是否成功处理该笔通知，false 则会重新通知
	Processed bool `json:"processed,omitempty"`
	// 处理结果
	Message string `json:"message,omitempty"`
}
