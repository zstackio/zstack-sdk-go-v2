// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CleanUpTrashOnPrimaryStorage operates on CleanUpTrashOnPrimaryStorage
func (cli *ZSClient) CleanUpTrashOnPrimaryStorage(uuid string, params param.CleanUpTrashOnPrimaryStorageParam) (*view.CleanUpTrashOnPrimaryStorageEventView, error) {
	resp := view.CleanUpTrashOnPrimaryStorageEventView{}
	if err := cli.Put("v1/primary-storage/{uuid}/trash/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
