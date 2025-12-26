// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateMonitorGroup updates MonitorGroup
func (cli *ZSClient) UpdateMonitorGroup(uuid string, params param.UpdateMonitorGroupParam) (*view.UpdateMonitorGroupEventView, error) {
	resp := view.UpdateMonitorGroupEventView{}
	if err := cli.Put("v1/zwatch/monitorgroups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
