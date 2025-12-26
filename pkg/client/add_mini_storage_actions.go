// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddMiniStorage adds MiniStorage
func (cli *ZSClient) AddMiniStorage(params param.AddMiniStorageParam) (*view.AddPrimaryStorageEventView, error) {
	resp := view.AddPrimaryStorageEventView{}
	if err := cli.Post("v1/primary-storage/mini", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
