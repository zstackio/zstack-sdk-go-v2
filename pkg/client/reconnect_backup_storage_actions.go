// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ReconnectBackupStorage operates on ReconnectBackupStorage
func (cli *ZSClient) ReconnectBackupStorage(uuid string, params param.ReconnectBackupStorageParam) (*view.ReconnectBackupStorageEventView, error) {
	resp := view.ReconnectBackupStorageEventView{}
	if err := cli.Put("v1/backup-storage/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
