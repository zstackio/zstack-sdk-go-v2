// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetPrimaryStorageCandidatesForVmMigration gets PrimaryStorageCandidatesForVmMigration by uuid
func (cli *ZSClient) GetPrimaryStorageCandidatesForVmMigration(uuid string) (*view.GetPrimaryStorageCandidatesForVmMigrationView, error) {
	var resp view.GetPrimaryStorageCandidatesForVmMigrationView
	if err := cli.Get("v1/vm-instances/{vmInstanceUuid}/storage-migration-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
