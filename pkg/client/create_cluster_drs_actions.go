// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateClusterDRS creates ClusterDRS
func (cli *ZSClient) CreateClusterDRS(params param.CreateClusterDRSParam) (*view.CreateClusterDRSEventView, error) {
	resp := view.CreateClusterDRSEventView{}
	if err := cli.Post("v1/clusters/{clusterUuid}/drs", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
