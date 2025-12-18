// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateMonitorGroup creates MonitorGroup
func (cli *ZSClient) CreateMonitorGroup(params param.CreateMonitorGroupParam) (*view.CreateMonitorGroupEventView, error) {
	resp := view.CreateMonitorGroupEventView{}
	if err := cli.Post("v1/zwatch/monitorgroups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
