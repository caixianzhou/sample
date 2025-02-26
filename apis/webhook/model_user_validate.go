package webhook

import "encoding/json"

type QueryUserInfoRequest struct {
	// 应用ID，由MidasBuy分配
	AppId string `json:"app_id,omitempty"`
	// 游戏玩家在游戏内的唯一标识
	UserId string `json:"user_id,omitempty"`
	// 游戏服务器的唯一标识符
	ServerId string `json:"server_id,omitempty"`
}

type QueryUserInfoResponse struct {
	// 应用ID，由MidasBuy分配
	AppId string `json:"app_id,omitempty"`
	// 用户ID
	UserId string `json:"user_id,omitempty"`
	// 游戏服务器的唯一标识符
	ServerId string `json:"server_id,omitempty"`
	// 用户名（游戏昵称）
	UserName string `json:"user_name,omitempty"`
	// 是否禁止用户充值。如果禁止，MidasBuy 会终止充值流程。如果不设置，缺省是不禁用 (可选)
	IsTopupForbidden bool `json:"is_topup_forbidden,omitempty"`
	// 禁止充值的原因 (可选)
	ForbiddenReason *ForbiddenReason `json:"forbidden_reason,omitempty"`
	// 一些充值流程需要的用户的扩展属性，根据具体场景来设置。如有需要请联系 Midas (可选)
	UserAttribute json.RawMessage `json:"user_attribute,omitempty"`
}

type ForbiddenReason struct {
	//禁止充值的错误码，业务自行定义，用于Midas定位的原因和统计。
	//不同业务的场景单独同步给 midasbuy
	//如：因为封号不满足: "1000001"
	//        因为未成年人的不满足："1000002"
	//        因为等级不足："1000003"
	Code string `json:"code,omitempty"`
	// 详情
	Message string `json:"message,omitempty"`
}
