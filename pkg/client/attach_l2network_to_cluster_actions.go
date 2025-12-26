// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachL2NetworkToCluster operates on L2NetworkToCluster
func (cli *ZSClient) AttachL2NetworkToCluster(params param.AttachL2NetworkToClusterParam) (*view.AttachL2NetworkToClusterEventView, error) {
	resp := view.AttachL2NetworkToClusterEventView{}
	if err := cli.Post("v1/l2-networks/{l2NetworkUuid}/clusters/{clusterUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
