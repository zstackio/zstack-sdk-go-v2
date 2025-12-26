// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetHostCandidatesForVmMigration gets HostCandidatesForVmMigration by uuid
func (cli *ZSClient) GetHostCandidatesForVmMigration(uuid string) (*view.GetHostCandidatesForVmMigrationView, error) {
	var resp view.GetHostCandidatesForVmMigrationView
	if err := cli.Get("v1/primary-storage/hosts/{vmInstanceUuid}/migration-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
