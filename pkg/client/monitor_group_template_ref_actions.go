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

func (cli *ZSClient) GetMonitorGroupTemplateRef(uuid string) (*view.MonitorGroupTemplateRefInventoryView, error) {
	var resp view.MonitorGroupTemplateRefInventoryView
	if err := cli.Get("v1/zwatch/monitorgroups/monitortemplates/refs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
