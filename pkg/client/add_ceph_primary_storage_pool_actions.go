// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddCephPrimaryStoragePool adds CephPrimaryStoragePool
func (cli *ZSClient) AddCephPrimaryStoragePool(params param.AddCephPrimaryStoragePoolParam) (*view.AddCephPrimaryStoragePoolEventView, error) {
	resp := view.AddCephPrimaryStoragePoolEventView{}
	if err := cli.Post("v1/primary-storage/ceph/{primaryStorageUuid}/pools", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
