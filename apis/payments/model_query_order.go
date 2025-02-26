package payments

import "git.woa.com/mbusiness/buy-api-library/midasbuy-go/apis/models"

type QueryPaymentOrderInfoRequest struct {
	// The ID of the app assigned by MidasBuy.
	Id string `json:"id,omitempty"`
}

type QueryPaymentOrderInfoResponse struct {
	// 应用ID
	AppId string `json:"app_id,omitempty"`
	// 交易发生时间，遵循 ISO 8601 标准的日期和时间表示格式，例如：2006-01-02T15:04:05Z07:00
	CreateTime string `json:"create_time,omitempty"`
	// 交易最后更新时间，遵循 ISO 8601 标准的日期和时间表示格式，例如：2006-01-02T15:04:05Z07:00
	UpdateTime string `json:"update_time,omitempty"`
	// 游戏玩家在游戏内的唯一标识
	UserId string `json:"user_id,omitempty"`
	// midasbuy 交易订单号，全局唯一
	PaymentOrderId string `json:"payment_order_id,omitempty"`
	// 服务器 id
	ServerId string `json:"server_id,omitempty"`
	// 订单状态
	OrderStatus models.OrderStatus `json:"order_status,omitempty"`
	// 物品列表
	OrderItems []*models.OrderItem `json:"order_items,omitempty"`
	// 交易总金额（单位元）
	TotalPrice *models.Price `json:"total_price,omitempty"`
	// 交易的站点国家地区
	ShopRegion string `json:"shop_region,omitempty"`
	// 支付渠道
	PaymentChannel string `json:"payment_channel,omitempty"`
}
