// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateCluster creates Cluster
func (cli *ZSClient) CreateCluster(params param.CreateClusterParam) (*view.CreateClusterEventView, error) {
	resp := view.CreateClusterEventView{}
	if err := cli.Post("v1/clusters", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
