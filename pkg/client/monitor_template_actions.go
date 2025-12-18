// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateMonitorTemplate 更新MonitorTemplate
func (cli *ZSClient) UpdateMonitorTemplate(uuid string, params param.UpdateMonitorTemplateParam) (*view.UpdateMonitorTemplateEventView, error) {
	resp := view.UpdateMonitorTemplateEventView{}
	if err := cli.Put("v1/zwatch/monitortemplates/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

