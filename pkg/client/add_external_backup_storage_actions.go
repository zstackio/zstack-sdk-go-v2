// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddExternalBackupStorage adds ExternalBackupStorage
func (cli *ZSClient) AddExternalBackupStorage(params param.AddExternalBackupStorageParam) (*view.AddExternalBackupStorageEventView, error) {
	resp := view.AddExternalBackupStorageEventView{}
	if err := cli.Post("v1/backup-storage/addon", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
