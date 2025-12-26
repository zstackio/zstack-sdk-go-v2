// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateWebhook updates Webhook
func (cli *ZSClient) UpdateWebhook(uuid string, params param.UpdateWebhookParam) (*view.UpdateWebhookEventView, error) {
	resp := view.UpdateWebhookEventView{}
	if err := cli.Put("v1/web-hooks/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
