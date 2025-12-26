// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetTrashOnBackupStorage gets TrashOnBackupStorage by uuid
func (cli *ZSClient) GetTrashOnBackupStorage(uuid string) (*view.GetTrashOnBackupStorageView, error) {
	var resp view.GetTrashOnBackupStorageView
	if err := cli.Get("v1/backup-storage/trash", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
