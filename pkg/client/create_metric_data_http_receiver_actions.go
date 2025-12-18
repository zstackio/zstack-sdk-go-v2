// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateMetricDataHttpReceiver creates MetricDataHttpReceiver
func (cli *ZSClient) CreateMetricDataHttpReceiver(params param.CreateMetricDataHttpReceiverParam) (*view.CreateMetricDataHttpReceiverEventView, error) {
	resp := view.CreateMetricDataHttpReceiverEventView{}
	if err := cli.Post("v1/zwatch/metrics/httpreceivers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
