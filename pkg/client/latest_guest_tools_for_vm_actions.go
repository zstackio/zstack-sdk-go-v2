// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetLatestGuestToolsForVm 获取LatestGuestToolsForVm详情
func (cli *ZSClient) GetLatestGuestToolsForVm(uuid string) (*view.GetLatestGuestToolsForVmView, error) {
	var resp view.GetLatestGuestToolsForVmView
	if err := cli.Get("v1/vm-instances/{uuid}/latest-guest-tools", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

