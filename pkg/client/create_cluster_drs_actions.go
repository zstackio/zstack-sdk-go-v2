// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateClusterDRS creates ClusterDRS
func (cli *ZSClient) CreateClusterDRS(params param.CreateClusterDRSParam) (*view.CreateClusterDRSEventView, error) {
	resp := view.CreateClusterDRSEventView{}
	if err := cli.Post("v1/clusters/{clusterUuid}/drs", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
