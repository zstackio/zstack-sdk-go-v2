// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetCandidateVmForAttachingIso gets CandidateVmForAttachingIso by uuid
func (cli *ZSClient) GetCandidateVmForAttachingIso(uuid string) (*view.GetCandidateVmForAttachingIsoView, error) {
	var resp view.GetCandidateVmForAttachingIsoView
	if err := cli.Get("v1/images/iso/{isoUuid}/vm-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
