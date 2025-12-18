// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetCandidateAffinityGroupForAttachingVm gets CandidateAffinityGroupForAttachingVm by uuid
func (cli *ZSClient) GetCandidateAffinityGroupForAttachingVm(uuid string) (*view.GetCandidateAffinityGroupForAttachingVmView, error) {
	var resp view.GetCandidateAffinityGroupForAttachingVmView
	if err := cli.Get("v1/affinityGroup/attachingVm", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
