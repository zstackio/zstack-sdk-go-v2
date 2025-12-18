// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetInterdependentL3NetworksBackupStorages 获取InterdependentL3NetworksBackupStorages详情
func (cli *ZSClient) GetInterdependentL3NetworksBackupStorages(uuid string) (*view.GetInterdependentL3NetworksBackupStoragesView, error) {
	var resp view.GetInterdependentL3NetworksBackupStoragesView
	if err := cli.Get("v1/backupStorage-l3networks/dependencies", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

