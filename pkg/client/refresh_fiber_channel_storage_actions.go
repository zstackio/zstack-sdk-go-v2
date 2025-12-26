// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RefreshFiberChannelStorage operates on RefreshFiberChannelStorage
func (cli *ZSClient) RefreshFiberChannelStorage(params param.RefreshFiberChannelStorageParam) (*view.RefreshFiberChannelStorageEventView, error) {
	resp := view.RefreshFiberChannelStorageEventView{}
	if err := cli.Post("v1/storage-devices/fiber-channel/controllers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
