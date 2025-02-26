package models

type Price struct {
	// 识别货币的三字符ISO-4217货币代码。见：https://en.wikipedia.org/wiki/ISO_4217#Active_codes
	Currency string `json:"currency,omitempty"`
	// 金额，单位为元。
	// 对于不同币种，我们有金额小数位的限制。比如：对于像JPY这样的货币，它通常不是分数，是一个整数，比如100。对于像TND这样的货币，它被细分为千分之一，是一个小数分数，比如100.123。对于货币代码所需的小数位数，请参见货币代码。
	Amount string `json:"amount,omitempty"`
}

type OrderItem struct {
	// 物品在游戏内的唯一标识
	ProductId string `json:"product_id,omitempty"`
	// 数量
	Quantity string `json:"quantity,omitempty"`
	// 商品类型
	ProductType ProductType `json:"product_type,omitempty"`
	// 交易总金额
	Price *Price `json:"price,omitempty"`
}

type PromotionItem struct {
	// 活动 ID
	PromotionId string `json:"promotion_id,omitempty"`
	// 模型 id
	ModelId string `json:"model_id,omitempty"`
	// 物品在游戏内的唯一标志
	ProductId string `json:"product_id,omitempty"`
	// 物品数量
	Quantity string `json:"quantity,omitempty"`
	// 物品类型
	ProductType ProductType `json:"product_type,omitempty"`
}
