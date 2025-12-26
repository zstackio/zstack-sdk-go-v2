// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateWebhook creates Webhook
func (cli *ZSClient) CreateWebhook(params param.CreateWebhookParam) (*view.CreateWebhookEventView, error) {
	resp := view.CreateWebhookEventView{}
	if err := cli.Post("v1/web-hooks", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
