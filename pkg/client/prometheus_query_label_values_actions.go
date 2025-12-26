// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// PrometheusQueryLabelValues operates on PrometheusQueryLabelValues
func (cli *ZSClient) PrometheusQueryLabelValues(params param.PrometheusQueryLabelValuesParam) (*view.PrometheusQueryLabelValuesView, error) {
	var resp view.PrometheusQueryLabelValuesView
	if err := cli.Get("v1/prometheus/labels", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
