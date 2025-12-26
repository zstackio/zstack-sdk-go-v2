// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SetIAM2ProjectContainerCluster operates on SetIAM2ProjectContainerCluster
func (cli *ZSClient) SetIAM2ProjectContainerCluster(uuid string, params param.SetIAM2ProjectContainerClusterParam) (*view.SetIAM2ProjectContainerClusterEventView, error) {
	resp := view.SetIAM2ProjectContainerClusterEventView{}
	if err := cli.Put("v1/iam2/projects/{uuid}/container/cluster/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
