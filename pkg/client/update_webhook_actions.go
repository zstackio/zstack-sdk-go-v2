// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateWebhook updates Webhook
func (cli *ZSClient) UpdateWebhook(uuid string, params param.UpdateWebhookParam) (*view.UpdateWebhookEventView, error) {
	resp := view.UpdateWebhookEventView{}
	if err := cli.Put("v1/web-hooks/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
