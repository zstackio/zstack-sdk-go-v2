// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SetVmBootMode 操作SetVmBootMode
func (cli *ZSClient) SetVmBootMode(uuid string, params param.SetVmBootModeParam) (*view.SetVmBootModeEventView, error) {
	resp := view.SetVmBootModeEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

