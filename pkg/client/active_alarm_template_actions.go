// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateActiveAlarmTemplate 更新ActiveAlarmTemplate
func (cli *ZSClient) UpdateActiveAlarmTemplate(uuid string, params param.UpdateActiveAlarmTemplateParam) (*view.UpdateActiveAlarmTemplateEventView, error) {
	resp := view.UpdateActiveAlarmTemplateEventView{}
	if err := cli.Put("v1/zwatch/activealarms/templates/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

