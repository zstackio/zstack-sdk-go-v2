// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateEmailMonitorTriggerAction updates EmailMonitorTriggerAction
func (cli *ZSClient) UpdateEmailMonitorTriggerAction(uuid string, params param.UpdateEmailMonitorTriggerActionParam) (*view.UpdateMonitorTriggerActionEventView, error) {
	resp := view.UpdateMonitorTriggerActionEventView{}
	if err := cli.Put("v1/monitoring/trigger-actions/emails/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
