// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CleanUpTrashOnBackupStorage operates on CleanUpTrashOnBackupStorage
func (cli *ZSClient) CleanUpTrashOnBackupStorage(uuid string, params param.CleanUpTrashOnBackupStorageParam) (*view.CleanUpTrashOnBackupStorageEventView, error) {
	resp := view.CleanUpTrashOnBackupStorageEventView{}
	if err := cli.Put("v1/backup-storage/{uuid}/trash/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
