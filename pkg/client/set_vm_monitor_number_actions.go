// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SetVmMonitorNumber 操作SetVmMonitorNumber
func (cli *ZSClient) SetVmMonitorNumber(uuid string, params param.SetVmMonitorNumberParam) (*view.SetVmMonitorNumberEventView, error) {
	resp := view.SetVmMonitorNumberEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

