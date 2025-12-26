// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateMonitorGroup creates MonitorGroup
func (cli *ZSClient) CreateMonitorGroup(params param.CreateMonitorGroupParam) (*view.CreateMonitorGroupEventView, error) {
	resp := view.CreateMonitorGroupEventView{}
	if err := cli.Post("v1/zwatch/monitorgroups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
