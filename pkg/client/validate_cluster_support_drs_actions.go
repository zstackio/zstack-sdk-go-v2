// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ValidateClusterSupportDRS operates on ValidateClusterSupportDRS
func (cli *ZSClient) ValidateClusterSupportDRS(params param.ValidateClusterSupportDRSParam) (*view.ValidateClusterSupportDRSView, error) {
	var resp view.ValidateClusterSupportDRSView
	if err := cli.Get("v1/clusters/{clusterUuid}/drs/valid", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
