// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// PrometheusQueryPassThrough operates on PrometheusQueryPassThrough
func (cli *ZSClient) PrometheusQueryPassThrough(params param.PrometheusQueryPassThroughParam) (*view.PrometheusQueryPassThroughView, error) {
	var resp view.PrometheusQueryPassThroughView
	if err := cli.Get("v1/prometheus/all", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
