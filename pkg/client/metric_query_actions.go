// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// GetAllMetricMetadata gets AllMetricMetadata by uuid
func (cli *ZSClient) GetAllMetricMetadata() (*view.GetAllMetricMetadataView, error) {
	var resp view.GetAllMetricMetadataView
	if err := cli.GetWithRespKey("v1/zwatch/metrics/meta-data", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetMetricData gets MetricData by uuid
func (cli *ZSClient) GetMetricData() (*view.GetMetricDataView, error) {
	var resp view.GetMetricDataView
	if err := cli.GetWithRespKey("v1/zwatch/metrics", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetMetricLabelValue gets MetricLabelValue by uuid
func (cli *ZSClient) GetMetricLabelValue() (*view.GetMetricLabelValueView, error) {
	var resp view.GetMetricLabelValueView
	if err := cli.GetWithRespKey("v1/zwatch/metrics/label-values", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
