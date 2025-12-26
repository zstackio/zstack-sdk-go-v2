// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetMetricLabelValue gets MetricLabelValue by uuid
func (cli *ZSClient) GetMetricLabelValue(uuid string) (*view.GetMetricLabelValueView, error) {
	var resp view.GetMetricLabelValueView
	if err := cli.Get("v1/zwatch/metrics/label-values", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
