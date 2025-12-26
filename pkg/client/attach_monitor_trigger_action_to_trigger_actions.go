// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachMonitorTriggerActionToTrigger operates on MonitorTriggerActionToTrigger
func (cli *ZSClient) AttachMonitorTriggerActionToTrigger(params param.AttachMonitorTriggerActionToTriggerParam) (*view.AttachMonitorTriggerActionToTriggerEventView, error) {
	resp := view.AttachMonitorTriggerActionToTriggerEventView{}
	if err := cli.Post("v1/monitoring/triggers/{triggerUuid}/trigger-actions/{actionUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
