// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetCandidateAffinityGroupForCreatingVm gets CandidateAffinityGroupForCreatingVm by uuid
func (cli *ZSClient) GetCandidateAffinityGroupForCreatingVm(uuid string) (*view.GetCandidateAffinityGroupForCreatingVmView, error) {
	var resp view.GetCandidateAffinityGroupForCreatingVmView
	if err := cli.Get("v1/vm-instances/candidate-affinityGroup", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
