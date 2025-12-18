// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateClusterOS 更新ClusterOS
func (cli *ZSClient) UpdateClusterOS(uuid string, params param.UpdateClusterOSParam) (*view.UpdateClusterOSEventView, error) {
	resp := view.UpdateClusterOSEventView{}
	if err := cli.Put("v1/clusters/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

