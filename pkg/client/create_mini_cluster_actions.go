// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateMiniCluster creates MiniCluster
func (cli *ZSClient) CreateMiniCluster(params param.CreateMiniClusterParam) (*view.CreateMiniClusterEventView, error) {
	resp := view.CreateMiniClusterEventView{}
	if err := cli.Post("v1/mini-clusters", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
