// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ApplyMonitorTemplateToMonitorGroup 操作ApplyMonitorTemplateToMonitorGroup
func (cli *ZSClient) ApplyMonitorTemplateToMonitorGroup(params param.ApplyMonitorTemplateToMonitorGroupParam) (*view.ApplyMonitorTemplateToMonitorGroupEventView, error) {
	resp := view.ApplyMonitorTemplateToMonitorGroupEventView{}
	if err := cli.Post("v1/zwatch/monitortemplates/{templateUuid}/monitorgroups/{groupUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

