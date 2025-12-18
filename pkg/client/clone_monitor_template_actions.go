// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CloneMonitorTemplate operates on MonitorTemplate
func (cli *ZSClient) CloneMonitorTemplate(params param.CloneMonitorTemplateParam) (*view.CloneMonitorTemplateEventView, error) {
	resp := view.CloneMonitorTemplateEventView{}
	if err := cli.Post("v1/zwatch/monitortemplates/{uuid}/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
