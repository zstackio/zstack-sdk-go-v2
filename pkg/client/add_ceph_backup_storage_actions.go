// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddCephBackupStorage adds CephBackupStorage
func (cli *ZSClient) AddCephBackupStorage(params param.AddCephBackupStorageParam) (*view.AddBackupStorageEventView, error) {
	resp := view.AddBackupStorageEventView{}
	if err := cli.Post("v1/backup-storage/ceph", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
