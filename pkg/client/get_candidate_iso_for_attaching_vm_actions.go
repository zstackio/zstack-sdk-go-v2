// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetCandidateIsoForAttachingVm gets CandidateIsoForAttachingVm by uuid
func (cli *ZSClient) GetCandidateIsoForAttachingVm(uuid string) (*view.GetCandidateIsoForAttachingVmView, error) {
	var resp view.GetCandidateIsoForAttachingVmView
	if err := cli.Get("v1/vm-instances/{vmInstanceUuid}/iso-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
