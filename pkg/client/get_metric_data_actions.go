// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetMetricData gets MetricData by uuid
func (cli *ZSClient) GetMetricData(uuid string) (*view.GetMetricDataView, error) {
	var resp view.GetMetricDataView
	if err := cli.Get("v1/zwatch/metrics", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
