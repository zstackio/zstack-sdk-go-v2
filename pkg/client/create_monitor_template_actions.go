// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateMonitorTemplate creates MonitorTemplate
func (cli *ZSClient) CreateMonitorTemplate(params param.CreateMonitorTemplateParam) (*view.CreateMonitorTemplateEventView, error) {
	resp := view.CreateMonitorTemplateEventView{}
	if err := cli.Post("v1/zwatch/monitortemplates", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
