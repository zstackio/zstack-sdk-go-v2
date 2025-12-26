// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetCandidateAffinityGroupForCreatingVm gets CandidateAffinityGroupForCreatingVm by uuid
func (cli *ZSClient) GetCandidateAffinityGroupForCreatingVm(uuid string) (*view.GetCandidateAffinityGroupForCreatingVmView, error) {
	var resp view.GetCandidateAffinityGroupForCreatingVmView
	if err := cli.Get("v1/vm-instances/candidate-affinityGroup", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
