// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryMonitorGroupEventSubscription queries MonitorGroupEventSubscription list
func (cli *ZSClient) QueryMonitorGroupEventSubscription(params *param.QueryParam) ([]view.MonitorGroupEventSubscriptionInventoryView, error) {
	var resp []view.MonitorGroupEventSubscriptionInventoryView
	return resp, cli.List("v1/zwatch/monitorgroups/subscriptions", params, &resp)
}

func (cli *ZSClient) GetMonitorGroupEventSubscription(uuid string) (*view.MonitorGroupEventSubscriptionInventoryView, error) {
	var resp view.MonitorGroupEventSubscriptionInventoryView
	if err := cli.Get("v1/zwatch/monitorgroups/subscriptions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageMonitorGroupEventSubscription Pagination
func (cli *ZSClient) PageMonitorGroupEventSubscription(params *param.QueryParam) ([]view.MonitorGroupEventSubscriptionInventoryView, int, error) {
	var monitorGroupEventSubscriptions []view.MonitorGroupEventSubscriptionInventoryView
	total, err := cli.Page("v1/zwatch/monitorgroups/subscriptions", params, &monitorGroupEventSubscriptions)
	return monitorGroupEventSubscriptions, total, err
}
