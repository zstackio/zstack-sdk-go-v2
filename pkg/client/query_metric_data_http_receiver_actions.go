// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryMetricDataHttpReceiver queries MetricDataHttpReceiver list
func (cli *ZSClient) QueryMetricDataHttpReceiver(params *param.QueryParam) ([]view.MetricDataHttpReceiverInventoryView, error) {
	var resp []view.MetricDataHttpReceiverInventoryView
	return resp, cli.List("v1/zwatch/metrics/httpreceivers", params, &resp)
}
