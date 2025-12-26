// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateEmailMonitorTriggerAction creates EmailMonitorTriggerAction
func (cli *ZSClient) CreateEmailMonitorTriggerAction(params param.CreateEmailMonitorTriggerActionParam) (*view.CreateMonitorTriggerActionEventView, error) {
	resp := view.CreateMonitorTriggerActionEventView{}
	if err := cli.Post("v1/monitoring/trigger-actions/emails", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
