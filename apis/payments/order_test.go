package payments

import (
	"context"
	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/core"
	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/core/consts"
	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/core/option"
	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/utils"
	"log"
	"strings"
	"testing"
)

const (
	// NOTE: 以下是随机生成的测试密钥，请勿用于生产环境
	testPrivateKey = "-----BEGIN TESTING KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDL1UzH2RAnJ6lU\nla0QLhoeBLmE63fxtvrFUMai+/KGzCHL29s4uKtmJgajT99M00kQEzf3+VFLXdfv\nub8rj1km9UtN3TdoNNgPPZ5EEnHZ8vzBdIVITWympbhCsL2Y+21lmZbZ3vLpu3Rb\nY6AzDnC37cIVxflZUwI5lpeDN4zob4KzOCWfHyRao23GCFTYn/UJ2edRkhP8FdNK\nItO7iiqutstjL+1WpcsNtLjeQ/wtXuI0N/qER/bmwdaoL9D/9WueoY56eoQarKB+\nouOSLl83dgYzZtNzYBAnGDRVw619cf+TIWzd8ji0cS59WevukYsFvd4cXJEg4bUp\nNnFt6uTNAgMBAAECggEAAY7m4Fw7cGEwPTJLuWTw1CvrEyYNq65famS8sABHEVq8\nI2fR3DQlM0m1IUh6B4dR9qp+8glY4r+b5/w+huG4p8CWS8kWJFjLEgrBi/msHyNp\nZT0zy6Kz4u4/Y1sgh+vcITu0WIQIzVqegBhZ4CoLGIzbv/jceB9XVANfsyQYkqpM\nEPpwY8mUYmJuQaXTb5n8gZujtzZjQpGkLeILS4L/LnbxFwC9nGU8ADygE7lO/nvU\ngrjuOSNANnS98Y2dk5h6HldqQqeAoJwZZPYNYYbGzzUB9U6Ei2l0aTY66eH5CnGE\n6RL6jmboWltWW2B4qU3Zh/VKRNiz8S95v4MSVGItQQKBgQDxoQ8kkvAVByApifEE\nNuUkNioWl0CL5mBT3KtkNzBYBc/RV3G/Tdmp1rx2jGPyGq+ETQaCI4OVndAAerW5\nA84TeigYH4PTV+XJ7c4ieQiO0bSnyxVM249IwdVhyvQLYglJlsq4LAbX9hWvNygX\noUjBhew4VdgFYPYfItUepPfbcQKBgQDX9Ma1TqBoxAW/2fDy54SY05y+aIJYEzRZ\nWOnNw5AksQPRtPaVsX/KnmDZZBLFMnSbcvnNKlx48JGRjkbt+hUYXOudJ3neLjjI\nvQ7UuwwX8ViIqHVexc3NspoitUfAjZT8Mo8aL2IMzt3xXkxhclMXRCA43gmCGSDC\n/1GQYXkZHQKBgQCpV/KP9HdUlXDiC+4hwQNpFJj8yjaPlf8O50ora05zcmdK1Vk/\n9STGllvxTcVCSZeXRpB4JsGy2y6LF3VC3LrSBbwR5Ax001aV5hehK2hnB+vv6THd\ncseB+288IYxWafgOXiNnXlvRgYODEEoF/aBLGTwL44YJhwIXokbxOjcH0QKBgQCN\nEQ4EPVo3VWTUD89/PJC3K/QFxUrvsYvOmXAQwyCTdzYhdG5nFk1907s8Bkzkl7Lo\nIFDhHjzNm4fbZu8aYPQKuBgIzlKjOdpJ9oWLnKunsDW+/xu8TsXDCln5NiWquFGL\n9JLZ7f3ElBUSqCCIvx9b4VqTCyd23mcyOYnUIHf0WQKBgQDPcIC9gd52QtOOZR9Q\nXIVbyWu3p9r3z6Po7F3dNtABb8smUwIrPaA5UiKPIpUfYmVmSwUCUK7y0HtrGB06\nfdl8giqDsN3pNVzs+LA2DnTURYkg823uN1Re5IpIdY9RAVylh6MUxaohJEueLmeD\n3RgAeinK3ZmO2ngBBbuV55VIJA==\n-----END TESTING KEY-----\n"
)

func TestOrderService_QueryPaymentOrder(t *testing.T) {
	var (
		mchID                      string = "test_mchID"         // 商户号
		mchCertificateSerialNumber string = "test_serial_number" // 商户证书序列号
	)

	// 使用 utils 提供的函数加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKey(testingKey(testPrivateKey))
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
			Id: "test_payment_order_id",
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

func testingKey(s string) string { return strings.ReplaceAll(s, "TESTING KEY", "PRIVATE KEY") }
