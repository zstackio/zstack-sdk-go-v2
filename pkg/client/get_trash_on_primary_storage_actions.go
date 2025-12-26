// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetTrashOnPrimaryStorage gets TrashOnPrimaryStorage by uuid
func (cli *ZSClient) GetTrashOnPrimaryStorage(uuid string) (*view.GetTrashOnPrimaryStorageView, error) {
	var resp view.GetTrashOnPrimaryStorageView
	if err := cli.Get("v1/primary-storage/trash", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
