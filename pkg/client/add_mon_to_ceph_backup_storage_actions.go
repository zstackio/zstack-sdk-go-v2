// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddMonToCephBackupStorage 操作AddMonToCephBackupStorage
func (cli *ZSClient) AddMonToCephBackupStorage(params param.AddMonToCephBackupStorageParam) (*view.AddMonToCephBackupStorageEventView, error) {
	resp := view.AddMonToCephBackupStorageEventView{}
	if err := cli.Post("v1/backup-storage/ceph/{uuid}/mons", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

