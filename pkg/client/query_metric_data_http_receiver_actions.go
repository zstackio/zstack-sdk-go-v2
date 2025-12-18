// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryMetricDataHttpReceiver queries MetricDataHttpReceiver list
func (cli *ZSClient) QueryMetricDataHttpReceiver(params param.QueryParam) ([]view.MetricDataHttpReceiverInventoryView, error) {
	var resp []view.MetricDataHttpReceiverInventoryView
	return resp, cli.List("v1/zwatch/metrics/httpreceivers", &params, &resp)
}
