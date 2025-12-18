// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachMonitorTriggerActionToTrigger 操作MonitorTriggerActionToTrigger
func (cli *ZSClient) AttachMonitorTriggerActionToTrigger(params param.AttachMonitorTriggerActionToTriggerParam) (*view.AttachMonitorTriggerActionToTriggerEventView, error) {
	resp := view.AttachMonitorTriggerActionToTriggerEventView{}
	if err := cli.Post("v1/monitoring/triggers/{triggerUuid}/trigger-actions/{actionUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

