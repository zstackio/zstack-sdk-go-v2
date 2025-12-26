// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateZBoxBackup creates ZBoxBackup
func (cli *ZSClient) CreateZBoxBackup(params param.CreateZBoxBackupParam) (*view.CreateExternalBackupEventView, error) {
	resp := view.CreateExternalBackupEventView{}
	if err := cli.Post("v1/externalbackup/zbox", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
