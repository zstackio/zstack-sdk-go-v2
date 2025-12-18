// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetCandidateIsoForAttachingVm gets CandidateIsoForAttachingVm by uuid
func (cli *ZSClient) GetCandidateIsoForAttachingVm(uuid string) (*view.GetCandidateIsoForAttachingVmView, error) {
	var resp view.GetCandidateIsoForAttachingVmView
	if err := cli.Get("v1/vm-instances/{vmInstanceUuid}/iso-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
