// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddExternalBackupStorage adds ExternalBackupStorage
func (cli *ZSClient) AddExternalBackupStorage(params param.AddExternalBackupStorageParam) (*view.AddExternalBackupStorageEventView, error) {
	resp := view.AddExternalBackupStorageEventView{}
	if err := cli.Post("v1/backup-storage/addon", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
