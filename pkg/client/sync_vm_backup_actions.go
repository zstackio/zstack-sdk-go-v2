// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SyncVmBackup 操作SyncVmBackup
func (cli *ZSClient) SyncVmBackup(uuid string, params param.SyncVmBackupParam) (*view.SyncVmBackupEventView, error) {
	resp := view.SyncVmBackupEventView{}
	if err := cli.Put("v1/vm-backups/imageStore/{imageStoreUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

