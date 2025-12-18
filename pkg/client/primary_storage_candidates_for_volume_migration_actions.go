// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetPrimaryStorageCandidatesForVolumeMigration 获取PrimaryStorageCandidatesForVolumeMigration详情
func (cli *ZSClient) GetPrimaryStorageCandidatesForVolumeMigration(uuid string) (*view.GetPrimaryStorageCandidatesForVolumeMigrationView, error) {
	var resp view.GetPrimaryStorageCandidatesForVolumeMigrationView
	if err := cli.Get("v1/primary-storage/volumes/{volumeUuid}/migration-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

