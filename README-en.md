# MidasBuy Go SDK

[MidasBuy] Official Go language client code library.

## Features

1. Interface SDK. See [apis](apis) 。
2. HTTP client `core.Client`, supports request signing and response verification. If the SDK does not support the interface you need, please use this client to initiate requests
3. Callback notification processing library `core/notify`, supports verification and decryption of MidasBuy callback notifications 

## Quick Start

### Installation

#### 1、 Use Go Modules to manage your project

If your project is not yet using Go Modules for dependency management, execute the following command in the root directory of the project:

```shell
go mod init
```

#### 2、 No need to clone the code in the repository, just execute the following command in the project directory

```shell
go get -u git.woa.com/mbusiness/buy-api-library/midasbuy-go
```

to add dependencies, complete the `go.mod` modification and SDK download.

### Sending Requests

First, initialize a `core.Client` instance, and then send requests to MidasBuy.

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
    var(
        mchID                      string="test_mch_id"// Merchant ID     
        mchCertificateSerialNumber string="test_serial_number"// Merchant certificate serial number
    ) 

    //Use the function provided by utils to load the merchant's private key, which will be used to generate the request signature 

    mchPrivateKey,err := utils.LoadPrivateKey("/path/to/merchant/apiclient_key.pem")
    if err!=nil {
        log.Print("load merchant private key error")
    }

    ctx := context.Background()

    opts := []core.ClientOption{
        option.WithMidasBuyPrivateKeyAuth(mchID,mchCertificateSerialNumber,mchPrivateKey),
    }

    client,err := core.NewClient(ctx,opts...)

    if err!=nil {
        log.Printf("new MidasBuy client err:%s",err)
    }

    svc := OrderService{
        Client:client,
        Env:consts.Sandbox,
    }

    resp,result,err := svc.QueryPaymentOrder(ctx, 
    QueryPaymentOrderInfoRequest{
        AppId:         mchID,
        PaymentOrderId:"test_payment_order_id",
        },
    )

    if err != nil{
        // Handle error 
        log.Printf("call QueryPaymentOrder err:%s",err)
    }else{
        // Handle response 
        log.Printf("status=%d resp=%v",result.Response.StatusCode,resp)
    }
} 
```

`resp` is the deserialized (UnmarshalJSON) response. In the above example, it is `*payments.QueryPaymentOrderInfoResponse` from the `apis/payments` package.

`result` is an instance of `*core.APIResult`, which contains the complete request message `*http.Request` and response message `*http.Response`.

#### Terminology

+ ** Merchant API Certificate **:   
Used to verify the merchant's identity. The certificate contains information such as the merchant ID, certificate serial number, and certificate validity period, issued by a Certificate Authority (CA) to prevent forgery or tampering.

+ ** Merchant API Private Key: **:   
When applying for a merchant API certificate, the merchant generates a private key, which is stored in the local certificate folder file apiclient_key.pem.

> :warning: Do not expose the private key file in public places, such as uploading it to Github, writing it in client code, etc.

+ ** MidasBuy Platform Certificate **:  
 The MidasBuy platform certificate is applied for by MidasBuy and contains the MidasBuy platform identifier and public key information. Merchants use the public key in the MidasBuy platform certificate to verify webhook signatures.

+ ** Certificate Serial Number **:
 Each certificate has a unique number issued by the CA, which is the certificate serial number.

## Error Handling

The SDK will return `error` in the following situations:

+ HTTP network errors, such as response reception timeout or network connection failure
+ Client failures, such as signature generation failure
+ The server returns a non `2xx` HTTP status code
+ Response signature verification failure

For convenience, the SDK converts `4xx` and `5xx` errors returned by the server into `APIError`.

```go
// Error handling example 
result,err := client.Get(ctx,string(consts.Sandbox))
if err!=nil{
    if core.IsAPIError(err,"INVALID_REQUEST"){
    // Handle invalid request 
    }
// Handle other errors      
}
```

## Verification of Callback Notifications

1. Use the MidasBuy platform certificate (for verification) to initialize `notify.Handler`
2. Call `handler.ParseNotifyRequest` to verify the signature and decrypt the message

### Initialization

+ Use the local MidasBuy platform public key to initialize `Handler` 。

Applicable scenario: First download the platform certificate to the local using the tool, and then use the locally managed platform certificate for verification and decryption.

```go
// 1. Initialize the MidasBuy platform certificate 
publicKey,err := utils.LoadPublicKeyWithPath("<midasbuy public key>")
// 2. Use publicKey to initialize `notify.Handler` 
handler := NewWebhookHandler(verifiers.NewSHA256WithRSAPubkeyVerifier(publicKeySerialNumber,*publicKey))
```

### Verification

Parse the content of the payment callback notification into `notify.Notification`.

```go
notifyRequest,err := handler.ParseNotifyRequest(request.Context(),request)
// If the verification fails or decryption fails 
if err != nil{
    fmt.Println(err)
    return
}

// Handle notification content, convert to specific instance according to EventType and ResourceType 
fmt.Println(notifyRequest.Id)
// For example, payment notification 
paymentNotification := PaymentNotification{}
err = json.Unmarshal(notification.Resource, &paymentNotification)
if err != nil {
    return err
}
fmt.Println(paymentNotification)

```

## FAQ

see [FAQ.md](FAQ.md)。

### Testing

The code submitted by developers should pass all test cases of this SDK.The SDK uses [agiledragon/gomonkey](https://github.com/agiledragon/gomonkey) and [stretchr/testify](https://github.com/stretchr/testify) in unit tests. Please confirm the relevant dependencies before testing. Use the following command to get all dependencies.

```bash
go get -t -v
```

Due to gomonkey, the parameter -gcflags=all=-l needs to be carried when executing test cases. Use the following command to initiate the test.

```bash
go test -gcflags=all=-l ./...
```
