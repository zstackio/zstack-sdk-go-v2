// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddNfsPrimaryStorage adds NfsPrimaryStorage
func (cli *ZSClient) AddNfsPrimaryStorage(params param.AddNfsPrimaryStorageParam) (*view.AddPrimaryStorageEventView, error) {
	resp := view.AddPrimaryStorageEventView{}
	if err := cli.Post("v1/primary-storage/nfs", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
