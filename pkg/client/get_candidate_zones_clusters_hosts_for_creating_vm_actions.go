// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetCandidateZonesClustersHostsForCreatingVm gets CandidateZonesClustersHostsForCreatingVm by uuid
func (cli *ZSClient) GetCandidateZonesClustersHostsForCreatingVm(uuid string) (*view.GetCandidateZonesClustersHostsForCreatingVmView, error) {
	var resp view.GetCandidateZonesClustersHostsForCreatingVmView
	if err := cli.Get("v1/vm-instances/candidate-destinations", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
