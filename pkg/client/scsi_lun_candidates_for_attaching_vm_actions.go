// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetScsiLunCandidatesForAttachingVm 获取ScsiLunCandidatesForAttachingVm详情
func (cli *ZSClient) GetScsiLunCandidatesForAttachingVm(uuid string) (*view.GetScsiLunCandidatesForAttachingVmView, error) {
	var resp view.GetScsiLunCandidatesForAttachingVmView
	if err := cli.Get("v1/vm-instances/{vmInstanceUuid}/candidate-storage-devices", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

