// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetCandidateAffinityGroupForAttachingVm gets CandidateAffinityGroupForAttachingVm by uuid
func (cli *ZSClient) GetCandidateAffinityGroupForAttachingVm(uuid string) (*view.GetCandidateAffinityGroupForAttachingVmView, error) {
	var resp view.GetCandidateAffinityGroupForAttachingVmView
	if err := cli.Get("v1/affinityGroup/attachingVm", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
