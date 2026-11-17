// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// GetPrometheusMetricLabelValue gets PrometheusMetricLabelValue by uuid
func (cli *ZSClient) GetPrometheusMetricLabelValue() (*view.GetPrometheusMetricLabelValueView, error) {
	var resp view.GetPrometheusMetricLabelValueView
	if err := cli.GetWithRespKey("v1/zwatch/metrics/prometheus/label-values", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PrometheusQueryLabelValues operates on PrometheusQueryLabelValues
func (cli *ZSClient) PrometheusQueryLabelValues() (*view.PrometheusQueryLabelValuesView, error) {
	var resp view.PrometheusQueryLabelValuesView
	if err := cli.GetWithRespKey("v1/prometheus/labels", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PrometheusQueryVmMonitoringData operates on PrometheusQueryVmMonitoringData
func (cli *ZSClient) PrometheusQueryVmMonitoringData() (*view.PrometheusQueryVmMonitoringDataView, error) {
	var resp view.PrometheusQueryVmMonitoringDataView
	if err := cli.GetWithRespKey("v1/prometheus/vm-instances", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PrometheusQueryMetadata operates on PrometheusQueryMetadata
func (cli *ZSClient) PrometheusQueryMetadata() (*view.PrometheusQueryMetadataView, error) {
	var resp view.PrometheusQueryMetadataView
	if err := cli.GetWithRespKey("v1/prometheus/meta-data", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PrometheusQueryPassThrough operates on PrometheusQueryPassThrough
func (cli *ZSClient) PrometheusQueryPassThrough() (*view.PrometheusQueryPassThroughView, error) {
	var resp view.PrometheusQueryPassThroughView
	if err := cli.GetWithRespKey("v1/prometheus/all", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
