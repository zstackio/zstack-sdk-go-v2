// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddLocalPrimaryStorage adds LocalPrimaryStorage
func (cli *ZSClient) AddLocalPrimaryStorage(params param.AddLocalPrimaryStorageParam) (*view.AddPrimaryStorageEventView, error) {
	resp := view.AddPrimaryStorageEventView{}
	if err := cli.Post("v1/primary-storage/local-storage", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
