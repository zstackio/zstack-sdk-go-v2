// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetCandidateVMForAttachingAffinityGroup gets CandidateVMForAttachingAffinityGroup by uuid
func (cli *ZSClient) GetCandidateVMForAttachingAffinityGroup(uuid string) (*view.GetCandidateVMForAttachingAffinityGroupView, error) {
	var resp view.GetCandidateVMForAttachingAffinityGroupView
	if err := cli.Get("v1/VM/attachingGroup", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
