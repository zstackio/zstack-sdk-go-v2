// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVmUsbRedirect gets VmUsbRedirect by uuid
func (cli *ZSClient) GetVmUsbRedirect(uuid string) (*view.GetVmUsbRedirectView, error) {
	var resp view.GetVmUsbRedirectView
	if err := cli.Get("v1/vm-instances/{uuid}/usbredirect", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
