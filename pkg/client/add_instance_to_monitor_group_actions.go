// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddInstanceToMonitorGroup 操作AddInstanceToMonitorGroup
func (cli *ZSClient) AddInstanceToMonitorGroup(params param.AddInstanceToMonitorGroupParam) (*view.AddInstanceToMonitorGroupEventView, error) {
	resp := view.AddInstanceToMonitorGroupEventView{}
	if err := cli.Post("v1/zwatch/monitorgroups/{groupUuid}/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

