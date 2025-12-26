// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ReconnectPrimaryStorage operates on ReconnectPrimaryStorage
func (cli *ZSClient) ReconnectPrimaryStorage(uuid string, params param.ReconnectPrimaryStorageParam) (*view.ReconnectPrimaryStorageEventView, error) {
	resp := view.ReconnectPrimaryStorageEventView{}
	if err := cli.Put("v1/primary-storage/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
