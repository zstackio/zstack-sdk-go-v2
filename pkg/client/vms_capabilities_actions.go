// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVmsCapabilities 获取VmsCapabilities详情
func (cli *ZSClient) GetVmsCapabilities(uuid string) (*view.GetVmsCapabilitiesView, error) {
	var resp view.GetVmsCapabilitiesView
	if err := cli.Get("v1/vm-instances/capabilities", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

