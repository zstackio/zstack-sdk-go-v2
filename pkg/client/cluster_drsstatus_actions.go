// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetClusterDRSStatus 获取ClusterDRSStatus详情
func (cli *ZSClient) GetClusterDRSStatus(uuid string) (*view.GetClusterDRSStatusView, error) {
	var resp view.GetClusterDRSStatusView
	if err := cli.Get("v1/clusters/drs/status", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

