// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetClusterHostNetworkFacts 获取ClusterHostNetworkFacts详情
func (cli *ZSClient) GetClusterHostNetworkFacts(uuid string) (*view.GetClusterHostNetworkFactsView, error) {
	var resp view.GetClusterHostNetworkFactsView
	if err := cli.Get("v1/cluster/hosts-network-facts/{clusterUuid}", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

