// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddInstanceToMonitorGroup adds InstanceToMonitorGroup
func (cli *ZSClient) AddInstanceToMonitorGroup(params param.AddInstanceToMonitorGroupParam) (*view.AddInstanceToMonitorGroupEventView, error) {
	resp := view.AddInstanceToMonitorGroupEventView{}
	if err := cli.Post("v1/zwatch/monitorgroups/{groupUuid}/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
