// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SetIAM2ProjectContainerCluster operates on SetIAM2ProjectContainerCluster
func (cli *ZSClient) SetIAM2ProjectContainerCluster(uuid string, params param.SetIAM2ProjectContainerClusterParam) (*view.SetIAM2ProjectContainerClusterEventView, error) {
	resp := view.SetIAM2ProjectContainerClusterEventView{}
	if err := cli.Put("v1/iam2/projects/{uuid}/container/cluster/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
