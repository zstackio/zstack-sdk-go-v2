// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetCandidateL2NetworksForAttachingCluster 获取CandidateL2NetworksForAttachingCluster详情
func (cli *ZSClient) GetCandidateL2NetworksForAttachingCluster(uuid string) (*view.GetCandidateL2NetworksForAttachingClusterView, error) {
	var resp view.GetCandidateL2NetworksForAttachingClusterView
	if err := cli.Get("v1/cluster/{clusterUuid}/l2-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

