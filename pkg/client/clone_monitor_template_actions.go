// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CloneMonitorTemplate operates on MonitorTemplate
func (cli *ZSClient) CloneMonitorTemplate(params param.CloneMonitorTemplateParam) (*view.CloneMonitorTemplateEventView, error) {
	resp := view.CloneMonitorTemplateEventView{}
	if err := cli.Post("v1/zwatch/monitortemplates/{uuid}/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
