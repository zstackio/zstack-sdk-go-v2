// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVmStartingCandidateClustersHosts gets VmStartingCandidateClustersHosts by uuid
func (cli *ZSClient) GetVmStartingCandidateClustersHosts(uuid string) (*view.GetVmStartingCandidateClustersHostsView, error) {
	var resp view.GetVmStartingCandidateClustersHostsView
	if err := cli.Get("v1/vm-instances/{uuid}/starting-target-hosts", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
