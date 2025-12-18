// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetCandidateVMForAttachingAffinityGroup 获取CandidateVMForAttachingAffinityGroup详情
func (cli *ZSClient) GetCandidateVMForAttachingAffinityGroup(uuid string) (*view.GetCandidateVMForAttachingAffinityGroupView, error) {
	var resp view.GetCandidateVMForAttachingAffinityGroupView
	if err := cli.Get("v1/VM/attachingGroup", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

