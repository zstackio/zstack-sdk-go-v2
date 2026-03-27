// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryMonitorGroupTemplateRef queries MonitorGroupTemplateRef list
func (cli *ZSClient) QueryMonitorGroupTemplateRef(ctx context.Context, params *param.QueryParam) ([]view.MonitorGroupTemplateRefInventoryView, error) {
	var resp []view.MonitorGroupTemplateRefInventoryView
	return resp, cli.List(ctx, "v1/zwatch/monitorgroups/monitortemplates/refs", params, &resp)
}

func (cli *ZSClient) GetMonitorGroupTemplateRef(ctx context.Context, uuid string) (*view.MonitorGroupTemplateRefInventoryView, error) {
	var resp view.MonitorGroupTemplateRefInventoryView
	if err := cli.Get(ctx, "v1/zwatch/monitorgroups/monitortemplates/refs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageMonitorGroupTemplateRef Pagination
func (cli *ZSClient) PageMonitorGroupTemplateRef(ctx context.Context, params *param.QueryParam) ([]view.MonitorGroupTemplateRefInventoryView, int, error) {
	var monitorGroupTemplateRefs []view.MonitorGroupTemplateRefInventoryView
	total, err := cli.Page(ctx, "v1/zwatch/monitorgroups/monitortemplates/refs", params, &monitorGroupTemplateRefs)
	return monitorGroupTemplateRefs, total, err
}
