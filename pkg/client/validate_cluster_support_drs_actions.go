// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ValidateClusterSupportDRS operates on ValidateClusterSupportDRS
func (cli *ZSClient) ValidateClusterSupportDRS(params param.ValidateClusterSupportDRSParam) (*view.ValidateClusterSupportDRSView, error) {
	var resp view.ValidateClusterSupportDRSView
	if err := cli.Get("v1/clusters/{clusterUuid}/drs/valid", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
