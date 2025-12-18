// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVmMigrationCandidateHosts 获取VmMigrationCandidateHosts详情
func (cli *ZSClient) GetVmMigrationCandidateHosts(uuid string) (*view.GetVmMigrationCandidateHostsView, error) {
	var resp view.GetVmMigrationCandidateHostsView
	if err := cli.Get("v1/vm-instances/{vmInstanceUuid}/migration-target-hosts", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

