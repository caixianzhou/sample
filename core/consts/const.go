package consts

import "time"

type Env string

// API 地址
const (
	Live    Env = "https://api-payments.midasbuy.com"
	Sandbox Env = "https://sandbox-api-payments.midasbuy.com"
)

// SDK 相关信息
const (
	Version         = "0.0.1"                     // SDK 版本
	UserAgentFormat = "MidasBuy-Go/%s (%s) GO/%s" // UserAgent中的信息
)

// HTTP 请求报文 Header 相关常量
const (
	Authorization = "Authorization"  // Header 中的 Authorization 字段
	Accept        = "Accept"         // Header 中的 Accept 字段
	ContentType   = "Content-Type"   // Header 中的 ContentType 字段
	ContentLength = "Content-Length" // Header 中的 ContentLength 字段
	UserAgent     = "User-Agent"     // Header 中的 UserAgent 字段
)

// 常用 ContentType
const (
	ApplicationJSON = "application/json"
	ImageJPG        = "image/jpg"
	ImagePNG        = "image/png"
	VideoMP4        = "video/mp4"
)

// 请求报文签名相关常量
const (
	SignatureMessageFormat = "%s\n%s\n%d\n%s\n%s\n" // 数字签名原文格式

	// HeaderAuthorizationPrefix Authorization前缀
	HeaderAuthorizationPrefix = "TXGW-"

	// HeaderAuthorizationFormat 请求头中的 Authorization 拼接格式
	HeaderAuthorizationFormat = "%s auth_id=\"%s\",auth_id_type=\"%s\",nonce_str=\"%s\",timestamp=\"%d\",serial_no=\"%s\"," +
		"signature=\"%s\"" // Authorization信息
)

// AuthIDType 鉴权类型
type AuthIDType string

// auth id type常量
const (
	// AuthIDTypeAppID 商户类型鉴权
	AuthIDTypeAppID AuthIDType = "app_id"
)

// HTTP 应答报文 Header 相关常量
const (
	MidasBuyTimestamp = "Txgw-Timestamp" // 回包时间戳
	MidasBuyNonce     = "Txgw-Nonce"     // 回包随机字符串
	MidasBuySignature = "Txgw-Signature" // 回包签名信息
	MidasBuySerial    = "Txgw-Serial"    // 回包平台序列号
	RequestID         = "Request-Id"     // 回包请求ID
)

// 时间相关常量
const (
	FiveMinute     = 5 * 60           // 回包校验最长时间（秒）
	DefaultTimeout = 30 * time.Second // HTTP 请求默认超时时间
)

// webhook相关资源
const (
	// ResourcePayment 支付通知
	ResourcePayment = "mpay.apis.event.PaymentNotification"
	// ResourceRefund 退款通知
	ResourceRefund = "mpay.apis.event.RefundNotification"
)

const (
	OrderQueryPath = "/midasbuy/v1/orders"
)
