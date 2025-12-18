// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeBackupStorageState 操作BackupStorageState
func (cli *ZSClient) ChangeBackupStorageState(uuid string, params param.ChangeBackupStorageStateParam) (*view.ChangeBackupStorageStateEventView, error) {
	resp := view.ChangeBackupStorageStateEventView{}
	if err := cli.Put("v1/backup-storage/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

