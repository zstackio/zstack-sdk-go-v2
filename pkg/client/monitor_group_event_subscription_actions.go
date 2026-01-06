// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryMonitorGroupEventSubscription queries MonitorGroupEventSubscription list
func (cli *ZSClient) QueryMonitorGroupEventSubscription(params *param.QueryParam) ([]view.MonitorGroupEventSubscriptionInventoryView, error) {
	var resp []view.MonitorGroupEventSubscriptionInventoryView
	return resp, cli.List("v1/zwatch/monitorgroups/subscriptions", params, &resp)
}
