// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ReconnectSftpBackupStorage operates on ReconnectSftpBackupStorage
func (cli *ZSClient) ReconnectSftpBackupStorage(uuid string, params param.ReconnectSftpBackupStorageParam) (*view.ReconnectSftpBackupStorageEventView, error) {
	resp := view.ReconnectSftpBackupStorageEventView{}
	if err := cli.Put("v1/backup-storage/sftp/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
