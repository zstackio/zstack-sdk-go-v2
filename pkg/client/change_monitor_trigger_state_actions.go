// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeMonitorTriggerState changes MonitorTriggerState
func (cli *ZSClient) ChangeMonitorTriggerState(uuid string, params param.ChangeMonitorTriggerStateParam) (*view.ChangeMonitorTriggerStateEventView, error) {
	resp := view.ChangeMonitorTriggerStateEventView{}
	if err := cli.Put("v1/monitoring/triggers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
