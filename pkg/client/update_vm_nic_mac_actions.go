// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateVmNicMac updates VmNicMac
func (cli *ZSClient) UpdateVmNicMac(uuid string, params param.UpdateVmNicMacParam) (*view.UpdateVmNicMacEventView, error) {
	resp := view.UpdateVmNicMacEventView{}
	if err := cli.Put("v1/vm-instances/nics/{vmNicUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
