// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateImageReplicationGroup 创建ImageReplicationGroup
func (cli *ZSClient) CreateImageReplicationGroup(params param.CreateImageReplicationGroupParam) (*view.CreateImageReplicationGroupEventView, error) {
	resp := view.CreateImageReplicationGroupEventView{}
	if err := cli.Post("v1/image-replication-groups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

