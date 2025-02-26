# MidasBuy Go SDK


[MidasBuy]官方Go语言客户端代码库。

## 功能介绍

1. 接口 SDK。详见 [接口介绍](apis)。
2. HTTP 客户端 `core.Client`，支持请求签名和应答验签。如果 SDK 未支持你需要的接口，请用此客户端发起请求。
3. 回调通知处理库 `core/notify`，支持MidasBuy回调通知的验签和解密。

## 快速开始

### 安装

#### 1、使用 Go Modules 管理你的项目

如果你的项目还不是使用 Go Modules 做依赖管理，在项目根目录下执行：

```shell
go mod init
```

#### 2、无需 clone 仓库中的代码，直接在项目目录中执行

```shell
go get -u git.woa.com/mbusiness/buy-api-library/midasbuy-go
```

来添加依赖，完成 `go.mod` 修改与 SDK 下载。

### 发送请求

先初始化一个 `core.Client` 实例，再向MidasBuy发送请求。

```go
package main

import (
	"context"
	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/core"
	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/core/consts"
	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/core/option"
	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/utils"
	"log"
)

func main() {
	var (
		mchID                      string = "test_mch_id"         // 商户号
		mchCertificateSerialNumber string = "test_serial_number" // 商户证书序列号
	)

	// 使用 utils 提供的函数加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKey("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Print("load merchant private key error")
	}

	ctx := context.Background()

	opts := []core.ClientOption{
		option.WithMidasBuyPrivateKeyAuth(mchID, mchCertificateSerialNumber, mchPrivateKey),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new MidasBuy client err:%s", err)
	}

	svc := OrderService{
		Client: client,
		Env:    consts.Sandbox,
	}
	resp, result, err := svc.QueryPaymentOrder(ctx,
		QueryPaymentOrderInfoRequest{
			AppId:          mchID,
			PaymentOrderId: "test_payment_order_id",
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call QueryPaymentOrder err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%v", result.Response.StatusCode, resp)
	}
}

```

`resp` 是反序列化（UnmarshalJSON）后的应答。上例中是 `apis/payments` 包中的 `*payments.QueryPaymentOrderInfoResponse`。

`result` 是 `*core.APIResult` 实例，包含了完整的请求报文 `*http.Request` 和应答报文 `*http.Response`。

#### 名词解释

+ **商户 API 证书**，是用来证实商户身份的。证书中包含商户号、证书序列号、证书有效期等信息，由证书授权机构（Certificate Authority ，简称 CA）签发，以防证书被伪造或篡改。

+ **商户 API 私钥**。商户申请商户 API 证书时，会生成商户私钥，并保存在本地证书文件夹的文件 apiclient_key.pem 中。

> :warning: 不要把私钥文件暴露在公共场合，如上传到 Github，写在客户端代码等。

+ **MidasBuy平台证书**。midasBuy平台证书是指由midasBuy负责申请的，包含midasBuy平台标识、公钥信息的证书。商户使用midasBuy平台证书中的公钥验证webhook签名。
+ **证书序列号**。每个证书都有一个由 CA 颁发的唯一编号，即证书序列号。
## 错误处理

以下情况，SDK 发送请求会返回 `error`：

+ HTTP 网络错误，如应答接收超时或网络连接失败
+ 客户端失败，如生成签名失败
+ 服务器端返回了**非** `2xx` HTTP 状态码
+ 应答签名验证失败

为了方便使用，SDK 将服务器返回的 `4xx` 和 `5xx` 错误，转换成了 `APIError`。

```go
// 错误处理示例
result, err := client.Get(ctx, string(consts.Sandbox))
if err != nil {
    if core.IsAPIError(err, "INVALID_REQUEST") { 
        // 处理无效请求 
    }
    // 处理的其他错误
}
```

## 回调通知的验签

1. 使用MidasBuy平台证书（验签）初始化 `notify.Handler`
2. 调用 `handler.ParseNotifyRequest` 验签，并解密报文。

### 初始化
+ 使用本地的midasBuy平台公钥初始化 `Handler`。

适用场景：首次通过工具下载平台证书到本地，后续使用本地管理的平台证书进行验签与解密。

```go
// 1. 初始化MidasBuy平台证书
publicKey, err := utils.LoadPublicKeyWithPath("<midasbuy public key>")
// 2. 使用publicKey初始化 `notify.Handler`
handler := NewWebhookHandler(verifiers.NewSHA256WithRSAPubkeyVerifier(publicKeySerialNumber, *publicKey))
```

### 验签

将支付回调通知中的内容，解析为 `notify.Notification`。

```go
notifyRequest, err := handler.ParseNotifyRequest(request.Context(), request)
// 如果验签未通过，或者解密失败
if err != nil {
    fmt.Println(err)
    return
}
// 处理通知内容, 根据EventType和ResourceType转为具体的实例
fmt.Println(notifyRequest.Id)
// 以支付通知为例
paymentNotification := PaymentNotification{}
err = json.Unmarshal(notification.Resource, &paymentNotification)
if err != nil {
    return err
}
fmt.Println(paymentNotification)
```

## 常见问题

常见问题请见 [FAQ.md](FAQ.md)。

### 测试

开发者提交的代码，应能通过本 SDK 所有的测试用例。

SDK 在单元测试中使用了 [agiledragon/gomonkey](https://github.com/agiledragon/gomonkey) 和 [stretchr/testify](https://github.com/stretchr/testify)，测试前请确认相关的依赖。使用以下命令获取所有的依赖。

```bash
go get -t -v
```

由于 `gomonkey` 的原因，在执行测试用例时需要携带参数 `-gcflags=all=-l`。使用以下命令发起测试。

```bash
go test -gcflags=all=-l ./...
```
