// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddMonToCephBackupStorage adds MonToCephBackupStorage
func (cli *ZSClient) AddMonToCephBackupStorage(params param.AddMonToCephBackupStorageParam) (*view.AddMonToCephBackupStorageEventView, error) {
	resp := view.AddMonToCephBackupStorageEventView{}
	if err := cli.Post("v1/backup-storage/ceph/{uuid}/mons", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
