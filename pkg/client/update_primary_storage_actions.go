// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdatePrimaryStorage updates PrimaryStorage
func (cli *ZSClient) UpdatePrimaryStorage(uuid string, params param.UpdatePrimaryStorageParam) (*view.UpdatePrimaryStorageEventView, error) {
	resp := view.UpdatePrimaryStorageEventView{}
	if err := cli.Put("v1/primary-storage/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
