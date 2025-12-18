// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateImageStoreBackupStorage 更新ImageStoreBackupStorage
func (cli *ZSClient) UpdateImageStoreBackupStorage(uuid string, params param.UpdateImageStoreBackupStorageParam) (*view.UpdateBackupStorageEventView, error) {
	resp := view.UpdateBackupStorageEventView{}
	if err := cli.Put("v1/backup-storage/image-store/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

