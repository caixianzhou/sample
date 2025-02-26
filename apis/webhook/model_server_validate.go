package webhook

import "git.woa.com/mbusiness/buy-api-library/midasbuy-go/apis/models"

type QueryServerInfoRequest struct {
	// 应用ID，由MidasBuy分配
	AppId string `json:"app_id,omitempty"`
	// ISO639标准的语言代码，例如en，见：https://en.wikipedia.org/wiki/ISO_639#Active_codes
	LanguageCode string `json:"language_code,omitempty"`
}

type QueryServerInfoResponse struct {
	// 应用ID，由MidasBuy分配
	AppId string `json:"app_id,omitempty"`
	// 游戏区服列表信息
	ServerItems []*ServerItem `json:"server_items,omitempty"`
}

type ServerItem struct {
	// 游戏服务器的唯一标识符
	ServerId string `json:"server_id,omitempty"`
	// 游戏服务器的名称
	ServerName string `json:"server_name,omitempty"`
	// 游戏服务器的状态
	ServerStatus models.ServerStatus `json:"server_status,omitempty"`
	// 区服分组，用于用户更方便的进行区服选择
	GroupName string `json:"group_name,omitempty"`
}
