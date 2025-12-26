// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVmGuestToolsInfo gets VmGuestToolsInfo by uuid
func (cli *ZSClient) GetVmGuestToolsInfo(uuid string) (*view.GetVmGuestToolsInfoView, error) {
	var resp view.GetVmGuestToolsInfoView
	if err := cli.Get("v1/vm-instances/{uuid}/guest-tools-infos", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
