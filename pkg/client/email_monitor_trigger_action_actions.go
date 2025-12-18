// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateEmailMonitorTriggerAction 更新EmailMonitorTriggerAction
func (cli *ZSClient) UpdateEmailMonitorTriggerAction(uuid string, params param.UpdateEmailMonitorTriggerActionParam) (*view.UpdateMonitorTriggerActionEventView, error) {
	resp := view.UpdateMonitorTriggerActionEventView{}
	if err := cli.Put("v1/monitoring/trigger-actions/emails/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

