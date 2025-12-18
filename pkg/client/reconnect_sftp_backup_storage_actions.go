// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ReconnectSftpBackupStorage 操作ReconnectSftpBackupStorage
func (cli *ZSClient) ReconnectSftpBackupStorage(uuid string, params param.ReconnectSftpBackupStorageParam) (*view.ReconnectSftpBackupStorageEventView, error) {
	resp := view.ReconnectSftpBackupStorageEventView{}
	if err := cli.Put("v1/backup-storage/sftp/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

