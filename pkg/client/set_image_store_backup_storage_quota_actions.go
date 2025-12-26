// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SetImageStoreBackupStorageQuota operates on SetImageStoreBackupStorageQuota
func (cli *ZSClient) SetImageStoreBackupStorageQuota(uuid string, params param.SetImageStoreBackupStorageQuotaParam) (*view.SetImageStoreBackupStorageQuotaEventView, error) {
	resp := view.SetImageStoreBackupStorageQuotaEventView{}
	if err := cli.Put("v1/backup-storage/image-store/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
