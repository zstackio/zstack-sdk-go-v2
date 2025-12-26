// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddSftpBackupStorage adds SftpBackupStorage
func (cli *ZSClient) AddSftpBackupStorage(params param.AddSftpBackupStorageParam) (*view.AddSftpBackupStorageEventView, error) {
	resp := view.AddSftpBackupStorageEventView{}
	if err := cli.Post("v1/backup-storage/sftp", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
