// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryMonitorGroupEventSubscription queries MonitorGroupEventSubscription list
func (cli *ZSClient) QueryMonitorGroupEventSubscription(params param.QueryParam) ([]view.MonitorGroupEventSubscriptionInventoryView, error) {
	var resp []view.MonitorGroupEventSubscriptionInventoryView
	return resp, cli.List("v1/zwatch/monitorgroups/subscriptions", &params, &resp)
}
