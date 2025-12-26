// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateClusterOS updates ClusterOS
func (cli *ZSClient) UpdateClusterOS(uuid string, params param.UpdateClusterOSParam) (*view.UpdateClusterOSEventView, error) {
	resp := view.UpdateClusterOSEventView{}
	if err := cli.Put("v1/clusters/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
