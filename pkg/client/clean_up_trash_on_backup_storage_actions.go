// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CleanUpTrashOnBackupStorage 操作CleanUpTrashOnBackupStorage
func (cli *ZSClient) CleanUpTrashOnBackupStorage(uuid string, params param.CleanUpTrashOnBackupStorageParam) (*view.CleanUpTrashOnBackupStorageEventView, error) {
	resp := view.CleanUpTrashOnBackupStorageEventView{}
	if err := cli.Put("v1/backup-storage/{uuid}/trash/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

