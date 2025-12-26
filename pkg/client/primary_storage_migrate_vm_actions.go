// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// PrimaryStorageMigrateVm operates on PrimaryStorageMigrateVm
func (cli *ZSClient) PrimaryStorageMigrateVm(uuid string, params param.PrimaryStorageMigrateVmParam) (*view.PrimaryStorageMigrateVmEventView, error) {
	resp := view.PrimaryStorageMigrateVmEventView{}
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
