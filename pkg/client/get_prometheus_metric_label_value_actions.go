// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetPrometheusMetricLabelValue gets PrometheusMetricLabelValue by uuid
func (cli *ZSClient) GetPrometheusMetricLabelValue(uuid string) (*view.GetPrometheusMetricLabelValueView, error) {
	var resp view.GetPrometheusMetricLabelValueView
	if err := cli.Get("v1/zwatch/metrics/prometheus/label-values", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
