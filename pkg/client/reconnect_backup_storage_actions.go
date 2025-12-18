// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ReconnectBackupStorage 操作ReconnectBackupStorage
func (cli *ZSClient) ReconnectBackupStorage(uuid string, params param.ReconnectBackupStorageParam) (*view.ReconnectBackupStorageEventView, error) {
	resp := view.ReconnectBackupStorageEventView{}
	if err := cli.Put("v1/backup-storage/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

