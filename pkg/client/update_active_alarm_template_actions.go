// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateActiveAlarmTemplate updates ActiveAlarmTemplate
func (cli *ZSClient) UpdateActiveAlarmTemplate(uuid string, params param.UpdateActiveAlarmTemplateParam) (*view.UpdateActiveAlarmTemplateEventView, error) {
	resp := view.UpdateActiveAlarmTemplateEventView{}
	if err := cli.Put("v1/zwatch/activealarms/templates/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
