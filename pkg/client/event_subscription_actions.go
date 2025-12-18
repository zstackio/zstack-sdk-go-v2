// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryEventSubscription 查询EventSubscription列表
func (cli *ZSClient) QueryEventSubscription(params param.QueryParam) ([]view.QueryEventSubscriptionView, error) {
	var resp []view.QueryEventSubscriptionView
	return resp, cli.List("v1/zwatch/events/subscriptions", &params, &resp)
}

