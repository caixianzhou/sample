// Copyright 2021 Inc. All rights reserved.

package notify_test

import (
	"context"
	"fmt"
	"net/http"

	"git.woa.com/mbusiness/buy-api-library/midasbuy-go/core/notify"
)

func ExampleHandler_ParseNotifyRequest_transaction() {
	var handler notify.Handler
	var request *http.Request

	content := make(map[string]interface{})
	notifyReq, err := handler.ParseNotifyRequest(context.Background(), request)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 处理通知内容
	fmt.Println(notifyReq.Summary)
	fmt.Println(content)
}
