// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateCluster updates Cluster
func (cli *ZSClient) UpdateCluster(uuid string, params param.UpdateClusterParam) (*view.UpdateClusterEventView, error) {
	resp := view.UpdateClusterEventView{}
	if err := cli.Put("v1/clusters/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
