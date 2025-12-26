// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RecoverDatabaseFromBackup operates on DatabaseFromBackup
func (cli *ZSClient) RecoverDatabaseFromBackup(uuid string, params param.RecoverDatabaseFromBackupParam) (*view.RecoverDatabaseFromBackupEventView, error) {
	resp := view.RecoverDatabaseFromBackupEventView{}
	if err := cli.Put("v1/database-backups/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
