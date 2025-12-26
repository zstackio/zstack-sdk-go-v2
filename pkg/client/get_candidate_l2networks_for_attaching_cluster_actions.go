// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetCandidateL2NetworksForAttachingCluster gets CandidateL2NetworksForAttachingCluster by uuid
func (cli *ZSClient) GetCandidateL2NetworksForAttachingCluster(uuid string) (*view.GetCandidateL2NetworksForAttachingClusterView, error) {
	var resp view.GetCandidateL2NetworksForAttachingClusterView
	if err := cli.Get("v1/cluster/{clusterUuid}/l2-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
