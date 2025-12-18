// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetCandidateClustersForAttachingL2Network gets CandidateClustersForAttachingL2Network by uuid
func (cli *ZSClient) GetCandidateClustersForAttachingL2Network(uuid string) (*view.GetCandidateClustersForAttachingL2NetworkView, error) {
	var resp view.GetCandidateClustersForAttachingL2NetworkView
	if err := cli.Get("v1/l2-networks/{l2NetworkUuid}/cluster-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
