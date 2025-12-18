// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// PrometheusQueryVmMonitoringData 操作PrometheusQueryVmMonitoringData
func (cli *ZSClient) PrometheusQueryVmMonitoringData(params param.PrometheusQueryVmMonitoringDataParam) (*view.PrometheusQueryVmMonitoringDataView, error) {
	var resp view.PrometheusQueryVmMonitoringDataView
	if err := cli.Get("v1/prometheus/vm-instances", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

