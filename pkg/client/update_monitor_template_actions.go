// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateMonitorTemplate updates MonitorTemplate
func (cli *ZSClient) UpdateMonitorTemplate(uuid string, params param.UpdateMonitorTemplateParam) (*view.UpdateMonitorTemplateEventView, error) {
	resp := view.UpdateMonitorTemplateEventView{}
	if err := cli.Put("v1/zwatch/monitortemplates/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
