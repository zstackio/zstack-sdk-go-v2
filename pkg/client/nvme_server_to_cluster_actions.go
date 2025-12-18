// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachNvmeServerToCluster 操作NvmeServerToCluster
func (cli *ZSClient) AttachNvmeServerToCluster(params param.AttachNvmeServerToClusterParam) (*view.AttachNvmeServerToClusterEventView, error) {
	resp := view.AttachNvmeServerToClusterEventView{}
	if err := cli.Post("v1/clusters/{clusterUuid}/storage-devices/nvme/servers/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

