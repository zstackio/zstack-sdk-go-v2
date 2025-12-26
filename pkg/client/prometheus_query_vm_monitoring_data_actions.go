// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// PrometheusQueryVmMonitoringData operates on PrometheusQueryVmMonitoringData
func (cli *ZSClient) PrometheusQueryVmMonitoringData(params param.PrometheusQueryVmMonitoringDataParam) (*view.PrometheusQueryVmMonitoringDataView, error) {
	var resp view.PrometheusQueryVmMonitoringDataView
	if err := cli.Get("v1/prometheus/vm-instances", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
