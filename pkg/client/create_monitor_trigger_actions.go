// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateMonitorTrigger creates MonitorTrigger
func (cli *ZSClient) CreateMonitorTrigger(params param.CreateMonitorTriggerParam) (*view.CreateMonitorTriggerEventView, error) {
	resp := view.CreateMonitorTriggerEventView{}
	if err := cli.Post("v1/monitoring/triggers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
