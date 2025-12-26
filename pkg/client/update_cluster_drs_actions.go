// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateClusterDRS updates ClusterDRS
func (cli *ZSClient) UpdateClusterDRS(uuid string, params param.UpdateClusterDRSParam) (*view.UpdateClusterDRSEventView, error) {
	resp := view.UpdateClusterDRSEventView{}
	if err := cli.Put("v1/clusters/drs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
