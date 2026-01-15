// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryMonitorGroupTemplateRef queries MonitorGroupTemplateRef list
func (cli *ZSClient) QueryMonitorGroupTemplateRef(params *param.QueryParam) ([]view.MonitorGroupTemplateRefInventoryView, error) {
	var resp []view.MonitorGroupTemplateRefInventoryView
	return resp, cli.List("v1/zwatch/monitorgroups/monitortemplates/refs", params, &resp)
}

// PageMonitorGroupTemplateRef Pagination
func (cli *ZSClient) PageMonitorGroupTemplateRef(params *param.QueryParam) ([]view.MonitorGroupTemplateRefInventoryView, int, error) {
	var monitorGroupTemplateRefs []view.MonitorGroupTemplateRefInventoryView
	total, err := cli.Page("v1/zwatch/monitorgroups/monitortemplates/refs", params, &monitorGroupTemplateRefs)
	return monitorGroupTemplateRefs, total, err
}
