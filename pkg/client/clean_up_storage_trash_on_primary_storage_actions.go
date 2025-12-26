// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CleanUpStorageTrashOnPrimaryStorage operates on CleanUpStorageTrashOnPrimaryStorage
func (cli *ZSClient) CleanUpStorageTrashOnPrimaryStorage(uuid string, params param.CleanUpStorageTrashOnPrimaryStorageParam) (*view.CleanUpStorageTrashOnPrimaryStorageEventView, error) {
	resp := view.CleanUpStorageTrashOnPrimaryStorageEventView{}
	if err := cli.Put("v1/primary-storage/{uuid}/storagetrash/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
