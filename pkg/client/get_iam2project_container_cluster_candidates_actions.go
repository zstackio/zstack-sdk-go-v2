// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetIAM2ProjectContainerClusterCandidates gets IAM2ProjectContainerClusterCandidates by uuid
func (cli *ZSClient) GetIAM2ProjectContainerClusterCandidates(uuid string) (*view.GetIAM2ProjectContainerClusterCandidatesView, error) {
	var resp view.GetIAM2ProjectContainerClusterCandidatesView
	if err := cli.Get("v1/iam2/projects/container/cluster/candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
