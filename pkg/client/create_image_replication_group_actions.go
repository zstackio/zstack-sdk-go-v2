// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateImageReplicationGroup creates ImageReplicationGroup
func (cli *ZSClient) CreateImageReplicationGroup(params param.CreateImageReplicationGroupParam) (*view.CreateImageReplicationGroupEventView, error) {
	resp := view.CreateImageReplicationGroupEventView{}
	if err := cli.Post("v1/image-replication-groups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
