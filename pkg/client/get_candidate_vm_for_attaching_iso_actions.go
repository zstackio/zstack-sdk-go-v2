// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetCandidateVmForAttachingIso gets CandidateVmForAttachingIso by uuid
func (cli *ZSClient) GetCandidateVmForAttachingIso(uuid string) (*view.GetCandidateVmForAttachingIsoView, error) {
	var resp view.GetCandidateVmForAttachingIsoView
	if err := cli.Get("v1/images/iso/{isoUuid}/vm-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
