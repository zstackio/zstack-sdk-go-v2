// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryEventSubscription queries EventSubscription list
func (cli *ZSClient) QueryEventSubscription(params param.QueryParam) ([]view.EventSubscriptionInventoryView, error) {
	var resp []view.EventSubscriptionInventoryView
	return resp, cli.List("v1/zwatch/events/subscriptions", &params, &resp)
}
