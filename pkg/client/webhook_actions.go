// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateWebhook 创建Webhook
func (cli *ZSClient) CreateWebhook(params param.CreateWebhookParam) (*view.CreateWebhookEventView, error) {
	resp := view.CreateWebhookEventView{}
	if err := cli.Post("v1/web-hooks", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

