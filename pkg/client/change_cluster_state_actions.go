// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeClusterState changes ClusterState
func (cli *ZSClient) ChangeClusterState(uuid string, params param.ChangeClusterStateParam) (*view.ChangeClusterStateEventView, error) {
	resp := view.ChangeClusterStateEventView{}
	if err := cli.Put("v1/clusters/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
