// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryWebhook queries Webhook list
func (cli *ZSClient) QueryWebhook(params *param.QueryParam) ([]view.WebhookInventoryView, error) {
	var resp []view.WebhookInventoryView
	return resp, cli.List("v1/web-hooks", params, &resp)
}
