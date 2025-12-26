// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ApplyMonitorTemplateToMonitorGroup operates on ApplyMonitorTemplateToMonitorGroup
func (cli *ZSClient) ApplyMonitorTemplateToMonitorGroup(params param.ApplyMonitorTemplateToMonitorGroupParam) (*view.ApplyMonitorTemplateToMonitorGroupEventView, error) {
	resp := view.ApplyMonitorTemplateToMonitorGroupEventView{}
	if err := cli.Post("v1/zwatch/monitortemplates/{templateUuid}/monitorgroups/{groupUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
