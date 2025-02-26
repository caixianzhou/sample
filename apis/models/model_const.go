package models

type ProductType string

const (
	VirtualCurrency ProductType = "VIRTUAL_CURRENCY" // 游戏币类型
	VirtualItem     ProductType = "VIRTUAL_ITEM"     // 道具类型
	Subscription    ProductType = "SUBSCRIPTION"     // 订阅类型，比如月卡/周卡
	RedeemCode      ProductType = "REDEEM_CODE"      // 兑换码
	Points          ProductType = "POINTS"           // 积分
	GrowthValue     ProductType = "GROWTH_VALUE"     //成长值
)

type OrderStatus string

const (
	Created                  OrderStatus = "CREATED"                    // 已下单
	Paid                     OrderStatus = "PAID"                       // 已支付
	Finished                 OrderStatus = "FINISHED"                   // 交易完成
	Cancelled                OrderStatus = "CANCELLED"                  // 已取消
	Refunding                OrderStatus = "REFUNDING"                  // 退款中
	RefundFailed             OrderStatus = "REFUND_FAILED"              // 退款失败
	Refunded                 OrderStatus = "REFUNDED"                   // 已经退款
	Chargeback               OrderStatus = "CHARGEBACK"                 // 拒付
	ChargebackReversed       OrderStatus = "CHARGEBACK_REVERSED"        // 拒付撤销
	SecondChargeback         OrderStatus = "SECOND_CHARGEBACK"          // 二次拒付
	SecondChargebackReversed OrderStatus = "SECOND_CHARGEBACK_REVERSED" // 二次拒付撤销
)

type ServerStatus string

const (
	ServerRunning ServerStatus = "SERVER_RUNNING" // 服务器正常
	ServerOffline ServerStatus = "SERVER_OFFLINE" // 服务器停服
)
