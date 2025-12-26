// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateCluster creates Cluster
func (cli *ZSClient) CreateCluster(params param.CreateClusterParam) (*view.CreateClusterEventView, error) {
	resp := view.CreateClusterEventView{}
	if err := cli.Post("v1/clusters", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
