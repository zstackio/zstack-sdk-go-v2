// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeMonitorTriggerActionState 操作MonitorTriggerActionState
func (cli *ZSClient) ChangeMonitorTriggerActionState(uuid string, params param.ChangeMonitorTriggerActionStateParam) (*view.ChangeMonitorTriggerActionStateEventView, error) {
	resp := view.ChangeMonitorTriggerActionStateEventView{}
	if err := cli.Put("v1/monitoring/trigger-actions/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

