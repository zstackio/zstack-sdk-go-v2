// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateMonitorTrigger updates MonitorTrigger
func (cli *ZSClient) UpdateMonitorTrigger(uuid string, params param.UpdateMonitorTriggerParam) (*view.UpdateMonitorTriggerEventView, error) {
	resp := view.UpdateMonitorTriggerEventView{}
	if err := cli.Put("v1/monitoring/triggers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
