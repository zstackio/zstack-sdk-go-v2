// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachNvmeServerToCluster operates on NvmeServerToCluster
func (cli *ZSClient) AttachNvmeServerToCluster(params param.AttachNvmeServerToClusterParam) (*view.AttachNvmeServerToClusterEventView, error) {
	resp := view.AttachNvmeServerToClusterEventView{}
	if err := cli.Post("v1/clusters/{clusterUuid}/storage-devices/nvme/servers/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
