// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SetImageStoreBackupStorageQuota 操作SetImageStoreBackupStorageQuota
func (cli *ZSClient) SetImageStoreBackupStorageQuota(uuid string, params param.SetImageStoreBackupStorageQuotaParam) (*view.SetImageStoreBackupStorageQuotaEventView, error) {
	resp := view.SetImageStoreBackupStorageQuotaEventView{}
	if err := cli.Put("v1/backup-storage/image-store/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

