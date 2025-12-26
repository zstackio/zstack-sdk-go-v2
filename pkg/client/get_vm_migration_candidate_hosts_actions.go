// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVmMigrationCandidateHosts gets VmMigrationCandidateHosts by uuid
func (cli *ZSClient) GetVmMigrationCandidateHosts(uuid string) (*view.GetVmMigrationCandidateHostsView, error) {
	var resp view.GetVmMigrationCandidateHostsView
	if err := cli.Get("v1/vm-instances/{vmInstanceUuid}/migration-target-hosts", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
