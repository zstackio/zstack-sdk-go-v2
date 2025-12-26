// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateVmNicDriver updates VmNicDriver
func (cli *ZSClient) UpdateVmNicDriver(uuid string, params param.UpdateVmNicDriverParam) (*view.UpdateVmNicDriverEventView, error) {
	resp := view.UpdateVmNicDriverEventView{}
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
