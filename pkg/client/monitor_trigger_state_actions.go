// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeMonitorTriggerState 操作MonitorTriggerState
func (cli *ZSClient) ChangeMonitorTriggerState(uuid string, params param.ChangeMonitorTriggerStateParam) (*view.ChangeMonitorTriggerStateEventView, error) {
	resp := view.ChangeMonitorTriggerStateEventView{}
	if err := cli.Put("v1/monitoring/triggers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

