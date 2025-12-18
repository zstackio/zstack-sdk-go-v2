// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetPrometheusMetricLabelValue 获取PrometheusMetricLabelValue详情
func (cli *ZSClient) GetPrometheusMetricLabelValue(uuid string) (*view.GetPrometheusMetricLabelValueView, error) {
	var resp view.GetPrometheusMetricLabelValueView
	if err := cli.Get("v1/zwatch/metrics/prometheus/label-values", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

