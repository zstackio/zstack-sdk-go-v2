// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddCephPrimaryStorage adds CephPrimaryStorage
func (cli *ZSClient) AddCephPrimaryStorage(params param.AddCephPrimaryStorageParam) (*view.AddPrimaryStorageEventView, error) {
	resp := view.AddPrimaryStorageEventView{}
	if err := cli.Post("v1/primary-storage/ceph", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
