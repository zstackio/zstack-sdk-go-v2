// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryMonitorGroupEventSubscription 查询MonitorGroupEventSubscription列表
func (cli *ZSClient) QueryMonitorGroupEventSubscription(params param.QueryParam) ([]view.QueryMonitorGroupEventSubscriptionView, error) {
	var resp []view.QueryMonitorGroupEventSubscriptionView
	return resp, cli.List("v1/zwatch/monitorgroups/subscriptions", &params, &resp)
}

