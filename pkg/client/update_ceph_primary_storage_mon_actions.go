// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateCephPrimaryStorageMon updates CephPrimaryStorageMon
func (cli *ZSClient) UpdateCephPrimaryStorageMon(uuid string, params param.UpdateCephPrimaryStorageMonParam) (*view.UpdateCephPrimaryStorageMonEventView, error) {
	resp := view.UpdateCephPrimaryStorageMonEventView{}
	if err := cli.Put("v1/primary-storage/ceph/mons/{monUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
