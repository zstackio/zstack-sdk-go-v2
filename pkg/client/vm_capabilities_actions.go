// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVmCapabilities 获取VmCapabilities详情
func (cli *ZSClient) GetVmCapabilities(uuid string) (*view.GetVmCapabilitiesView, error) {
	var resp view.GetVmCapabilitiesView
	if err := cli.Get("v1/vm-instances/{uuid}/capabilities", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

