// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryMonitorGroupInstance queries MonitorGroupInstance list
func (cli *ZSClient) QueryMonitorGroupInstance(params *param.QueryParam) ([]view.MonitorGroupInstanceInventoryView, error) {
	var resp []view.MonitorGroupInstanceInventoryView
	return resp, cli.List("v1/zwatch/monitorgroups/instances", params, &resp)
}

// PageMonitorGroupInstance Pagination
func (cli *ZSClient) PageMonitorGroupInstance(params *param.QueryParam) ([]view.MonitorGroupInstanceInventoryView, int, error) {
	var monitorGroupInstances []view.MonitorGroupInstanceInventoryView
	total, err := cli.Page("v1/zwatch/monitorgroups/instances", params, &monitorGroupInstances)
	return monitorGroupInstances, total, err
}
