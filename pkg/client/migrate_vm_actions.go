// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// MigrateVm operates on Vm
func (cli *ZSClient) MigrateVm(uuid string, params param.MigrateVmParam) (*view.MigrateVmEventView, error) {
	resp := view.MigrateVmEventView{}
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
