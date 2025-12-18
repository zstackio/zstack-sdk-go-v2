// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVmConsoleAddress 获取VmConsoleAddress详情
func (cli *ZSClient) GetVmConsoleAddress(uuid string) (*view.GetVmConsoleAddressView, error) {
	var resp view.GetVmConsoleAddressView
	if err := cli.Get("v1/vm-instances/{uuid}/console-addresses", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

