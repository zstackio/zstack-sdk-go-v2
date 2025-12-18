// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddResourceStackVmPortMonitor 操作AddResourceStackVmPortMonitor
func (cli *ZSClient) AddResourceStackVmPortMonitor(params param.AddResourceStackVmPortMonitorParam) (*view.AddResourceStackVmPortMonitorEventView, error) {
	resp := view.AddResourceStackVmPortMonitorEventView{}
	if err := cli.Post("v1/cloudformation/stack/monitor/addvm", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

