package webhook

import (
	"context"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/apis/models"
	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/core/auth/verifiers"
	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/core/consts"
	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/core/notify"
	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/utils"
	"log"
	"net/http"
	"testing"
)

var (
	url       = "/your_webhook_url"
	publicKey *rsa.PublicKey
	handler   *notify.Handler
	err       error
)

func TestNewWebhookHandler(t *testing.T) {
	WebHookServer()
}

func WebHookServer() {
	// 使用 utils 提供的函数从本地文件中加载验签公钥，公钥会用来验证midasBuy的webhook的有效性
	publicKey, err = utils.LoadPublicKeyWithPath("/path/to/midas_buy/public_key.pem")
	if err != nil {
		log.Print("load merchant private key error")
	}
	handler = NewWebhookHandler(verifiers.NewSHA256WithRSAPubkeyVerifier(*publicKey))
	mux := http.NewServeMux()
	mux.HandleFunc(url, func(writer http.ResponseWriter, request *http.Request) {
		var (
			err          error
			responseBody []byte
		)
		defer func() {
			writer.Header().Set(consts.ContentType, consts.ApplicationJSON)
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
			} else {
				writer.WriteHeader(http.StatusOK)
				_, _ = writer.Write(responseBody)
			}
		}()
		notifyRequest, err := handler.ParseNotifyRequest(request.Context(), request)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("webhook: %#v\n", notifyRequest)
		switch {
		case notifyRequest.EventType == notify.PaymentOrderStatusUpdate &&
			notifyRequest.ResourceType == notify.ResourceTypeOrder:
			responseBody, err = paymentNotify(request.Context(), notifyRequest)
			if err != nil {
				return
			}
		case notifyRequest.EventType == notify.PromotionOrderFinished &&
			notifyRequest.ResourceType == notify.ResourceTypePromotion:
			responseBody, err = promotionNotify(request.Context(), notifyRequest)
			if err != nil {
				return
			}
		case notifyRequest.EventType == notify.ServerValidate &&
			notifyRequest.ResourceType == notify.ResourceTypeServer:
			responseBody, err = serverValidate(request.Context(), notifyRequest)
			if err != nil {
				return
			}
		case notifyRequest.EventType == notify.UserValidate &&
			notifyRequest.ResourceType == notify.ResourceTypeUser:
			responseBody, err = userValidate(request.Context(), notifyRequest)
			if err != nil {
				return
			}
		case notifyRequest.EventType == notify.ProductValidate &&
			notifyRequest.ResourceType == notify.ResourceTypeProduct:
			responseBody, err = productValidate(request.Context(), notifyRequest)
			if err != nil {
				return
			}
		default:
			fmt.Printf("unknown Event: %s and Resource: %s", notifyRequest.EventType, notifyRequest.ResourceType)
		}
		return
	})
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}

}

func userValidate(ctx context.Context, notification *notify.Notification) (responseBody []byte, err error) {
	userInfoRequest := QueryUserInfoRequest{}
	err = json.Unmarshal(notification.Resource, &userInfoRequest)
	if err != nil {
		return nil, err
	}
	//todo 根据userInfoRequest组装response
	response := QueryUserInfoResponse{
		AppId:            userInfoRequest.AppId,
		UserId:           userInfoRequest.UserId,
		ServerId:         "1",
		UserName:         "test user",
		IsTopupForbidden: false,
		ForbiddenReason: &ForbiddenReason{
			Code:    "",
			Message: "",
		},
		UserAttribute: nil,
	}
	//
	responseBody, err = json.Marshal(response)
	if err != nil {
		return nil, err
	}
	return responseBody, nil
}

func productValidate(ctx context.Context, notification *notify.Notification) (responseBody []byte, err error) {
	userPurchaseEligibilityRequest := CheckUserPurchaseEligibilityRequest{}
	err = json.Unmarshal(notification.Resource, &userPurchaseEligibilityRequest)
	if err != nil {
		return nil, err
	}
	// todo 根据userPurchaseEligibilityRequest组装response
	response := CheckUserPurchaseEligibilityResponse{
		ProductEligibleInfos: make([]*ProductEligibleInfo, len(userPurchaseEligibilityRequest.ProductItems)),
	}
	for _, v := range userPurchaseEligibilityRequest.ProductItems {
		productEligibleInfo := ProductEligibleInfo{
			ProductId:           v.ProductId,
			IsPurchaseForbidden: false,
			Code:                "",
			Message:             "",
			TimeToNextPurchase:  0,
		}
		response.ProductEligibleInfos = append(response.ProductEligibleInfos, &productEligibleInfo)
	}
	//
	responseBody, err = json.Marshal(response)
	responseBody, err = json.Marshal(response)
	if err != nil {
		return nil, err
	}
	return responseBody, nil
}

func serverValidate(ctx context.Context, notification *notify.Notification) (responseBody []byte, err error) {
	serverInfoRequest := QueryServerInfoRequest{}
	err = json.Unmarshal(notification.Resource, &serverInfoRequest)
	if err != nil {
		return nil, err
	}
	// todo 根据serverInfoRequest组装response
	response := QueryServerInfoResponse{
		AppId: serverInfoRequest.AppId,
		ServerItems: []*ServerItem{
			{
				ServerId:     "1",
				ServerName:   "test",
				ServerStatus: models.ServerRunning,
				GroupName:    "group test",
			},
		},
	}
	//
	responseBody, err = json.Marshal(response)
	if err != nil {
		return nil, err
	}
	return responseBody, nil
}

func paymentNotify(ctx context.Context, notification *notify.Notification) (responseBody []byte, err error) {
	paymentNotification := PaymentNotification{}
	err = json.Unmarshal(notification.Resource, &paymentNotification)
	if err != nil {
		return nil, err
	}
	// todo 根据paymentNotification组装response
	response := NotifyResponse{
		Processed: false,
		Message:   "",
	}
	// todo 游戏需进行幂等验证,防止重复发货
	if mockGameOrderProvided(ctx, paymentNotification.PaymentOrderId) {
		_ = paymentNotification
		response.Processed = true
	} else {
		// todo 游戏自身逻辑处理
		_ = paymentNotification
		response.Processed = true
	}
	//
	responseBody, err = json.Marshal(response)
	if err != nil {
		return nil, err
	}
	return responseBody, nil
}

func promotionNotify(ctx context.Context, notification *notify.Notification) (responseBody []byte, err error) {
	promotionNotification := PromotionNotification{}
	err = json.Unmarshal(notification.Resource, &promotionNotification)
	if err != nil {
		return nil, err
	}
	// todo 根据promotionNotification组装response
	response := NotifyResponse{
		Processed: false,
		Message:   "",
	}
	// todo 游戏需进行幂等验证,防止重复发货
	if mockGameOrderProvided(ctx, promotionNotification.PaymentOrderId) {
		_ = promotionNotification
		response.Processed = true
	} else {
		// todo 游戏自身逻辑处理
		_ = promotionNotification
		response.Processed = true
	}
	//
	responseBody, err = json.Marshal(response)
	if err != nil {
		return nil, err
	}
	return responseBody, nil
}

func mockGameOrderProvided(ctx context.Context, paymentOrderId string) bool {
	return true
}
