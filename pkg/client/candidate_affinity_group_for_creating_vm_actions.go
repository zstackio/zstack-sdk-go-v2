// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetCandidateAffinityGroupForCreatingVm 获取CandidateAffinityGroupForCreatingVm详情
func (cli *ZSClient) GetCandidateAffinityGroupForCreatingVm(uuid string) (*view.GetCandidateAffinityGroupForCreatingVmView, error) {
	var resp view.GetCandidateAffinityGroupForCreatingVmView
	if err := cli.Get("v1/vm-instances/candidate-affinityGroup", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

