// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateClusterDRS 更新ClusterDRS
func (cli *ZSClient) UpdateClusterDRS(uuid string, params param.UpdateClusterDRSParam) (*view.UpdateClusterDRSEventView, error) {
	resp := view.UpdateClusterDRSEventView{}
	if err := cli.Put("v1/clusters/drs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

