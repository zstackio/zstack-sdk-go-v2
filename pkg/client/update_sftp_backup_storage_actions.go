// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateSftpBackupStorage updates SftpBackupStorage
func (cli *ZSClient) UpdateSftpBackupStorage(uuid string, params param.UpdateSftpBackupStorageParam) (*view.UpdateBackupStorageEventView, error) {
	resp := view.UpdateBackupStorageEventView{}
	if err := cli.Put("v1/backup-storage/sftp/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
