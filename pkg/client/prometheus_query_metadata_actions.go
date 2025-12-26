// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// PrometheusQueryMetadata operates on PrometheusQueryMetadata
func (cli *ZSClient) PrometheusQueryMetadata(params param.PrometheusQueryMetadataParam) (*view.PrometheusQueryMetadataView, error) {
	var resp view.PrometheusQueryMetadataView
	if err := cli.Get("v1/prometheus/meta-data", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
