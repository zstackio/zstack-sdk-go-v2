// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetScsiLunCandidatesForAttachingVm gets ScsiLunCandidatesForAttachingVm by uuid
func (cli *ZSClient) GetScsiLunCandidatesForAttachingVm(uuid string) (*view.GetScsiLunCandidatesForAttachingVmView, error) {
	var resp view.GetScsiLunCandidatesForAttachingVmView
	if err := cli.Get("v1/vm-instances/{vmInstanceUuid}/candidate-storage-devices", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
