// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ReconnectImageStoreBackupStorage 操作ReconnectImageStoreBackupStorage
func (cli *ZSClient) ReconnectImageStoreBackupStorage(uuid string, params param.ReconnectImageStoreBackupStorageParam) (*view.ReconnectImageStoreBackupStorageEventView, error) {
	resp := view.ReconnectImageStoreBackupStorageEventView{}
	if err := cli.Put("v1/backup-storage/image-store/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

