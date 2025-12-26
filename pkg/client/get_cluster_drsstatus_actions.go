// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetClusterDRSStatus gets ClusterDRSStatus by uuid
func (cli *ZSClient) GetClusterDRSStatus(uuid string) (*view.GetClusterDRSStatusView, error) {
	var resp view.GetClusterDRSStatusView
	if err := cli.Get("v1/clusters/drs/status", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
