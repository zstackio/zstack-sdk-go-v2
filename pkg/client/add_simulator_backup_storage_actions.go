// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddSimulatorBackupStorage 操作AddSimulatorBackupStorage
func (cli *ZSClient) AddSimulatorBackupStorage(params param.AddSimulatorBackupStorageParam) (*view.AddBackupStorageEventView, error) {
	resp := view.AddBackupStorageEventView{}
	if err := cli.Post("v1/backup-storage/simulators", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

