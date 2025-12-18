// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateEcsImageFromEcsSnapshot creates EcsImageFromEcsSnapshot
func (cli *ZSClient) CreateEcsImageFromEcsSnapshot(params param.CreateEcsImageFromEcsSnapshotParam) (*view.CreateEcsImageFromEcsSnapshotEventView, error) {
	resp := view.CreateEcsImageFromEcsSnapshotEventView{}
	if err := cli.Post("v1/hybrid/aliyun/image/snapshot", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
