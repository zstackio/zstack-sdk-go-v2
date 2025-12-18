// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachBaremetalPxeServerToCluster 操作BaremetalPxeServerToCluster
func (cli *ZSClient) AttachBaremetalPxeServerToCluster(params param.AttachBaremetalPxeServerToClusterParam) (*view.AttachBaremetalPxeServerToClusterEventView, error) {
	resp := view.AttachBaremetalPxeServerToClusterEventView{}
	if err := cli.Post("v1/clusters/{clusterUuid}/pxeservers/{pxeServerUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

