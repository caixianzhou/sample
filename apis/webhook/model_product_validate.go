package webhook

import "git.woa.com/mbusiness/buy-api-library/midasbuy-go/apis/models"

type CheckUserPurchaseEligibilityRequest struct {
	// 应用ID，由MidasBuy分配
	AppId string `json:"app_id,omitempty"`
	// 用户ID
	UserId string `json:"user_id,omitempty"`
	// 游戏服务器的唯一标识符
	ServerId string `json:"server_id,omitempty"`
	// 物品列表
	ProductItems []*ProductItem `json:"product_items,omitempty"`
	// midas站点
	ShopRegion string `json:"shop_region,omitempty"`
}

type CheckUserPurchaseEligibilityResponse struct {
	// 查询资格应答，如果未传回查询中的 product_id则默认可购买
	ProductEligibleInfos []*ProductEligibleInfo `json:"product_eligible_infos,omitempty"`
}

type ProductEligibleInfo struct {
	// 物品 id
	ProductId string `json:"product_id,omitempty"`
	// 是否禁止购买，为 true 表示禁止购买
	IsPurchaseForbidden bool `json:"is_purchase_forbidden,omitempty"`
	// 禁止原因
	// 没有资格的错误码，用于定位不满足资格的原因。
	// 不同业务的场景单独同步给 midasbuy
	// 如：因为用户已经拥有: "1000001"
	//    因为封号导致的不满足："1000002"
	//    因为用户未满足购买资格："1000003"
	Code string `json:"code,omitempty"`
	// 详情
	Message string `json:"message,omitempty"`
	// 距离下次可购买需要等待时间，单位秒
	TimeToNextPurchase int64 `json:"time_to_next_purchase,omitempty"`
}

type ProductItem struct {
	// 物品在游戏内的唯一标识
	ProductId string `json:"product_id,omitempty"`
	// 物品数量
	Quantity string `json:"quantity,omitempty"`
	// 物品类型
	ProductType models.ProductType `json:"product_type,omitempty"`
}
