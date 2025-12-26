// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachBaremetalPxeServerToCluster operates on BaremetalPxeServerToCluster
func (cli *ZSClient) AttachBaremetalPxeServerToCluster(params param.AttachBaremetalPxeServerToClusterParam) (*view.AttachBaremetalPxeServerToClusterEventView, error) {
	resp := view.AttachBaremetalPxeServerToClusterEventView{}
	if err := cli.Post("v1/clusters/{clusterUuid}/pxeservers/{pxeServerUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
