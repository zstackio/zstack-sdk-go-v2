// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetPrimaryStorageCandidatesForVmMigration 获取PrimaryStorageCandidatesForVmMigration详情
func (cli *ZSClient) GetPrimaryStorageCandidatesForVmMigration(uuid string) (*view.GetPrimaryStorageCandidatesForVmMigrationView, error) {
	var resp view.GetPrimaryStorageCandidatesForVmMigrationView
	if err := cli.Get("v1/vm-instances/{vmInstanceUuid}/storage-migration-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

