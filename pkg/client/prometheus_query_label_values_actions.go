// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// PrometheusQueryLabelValues 操作PrometheusQueryLabelValues
func (cli *ZSClient) PrometheusQueryLabelValues(params param.PrometheusQueryLabelValuesParam) (*view.PrometheusQueryLabelValuesView, error) {
	var resp view.PrometheusQueryLabelValuesView
	if err := cli.Get("v1/prometheus/labels", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

