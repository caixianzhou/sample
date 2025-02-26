package webhook

import (
	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/core/auth"
	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/core/auth/validators"
	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/core/notify"
)

// NewWebhookHandler 创建通知处理器
func NewWebhookHandler(
	verifier auth.Verifier,
) *notify.Handler {
	return &notify.Handler{
		Validator: *validators.NewMidasBuyNotifyValidator(verifier),
	}
}
