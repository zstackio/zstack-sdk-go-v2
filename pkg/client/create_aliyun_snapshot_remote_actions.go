// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateAliyunSnapshotRemote creates AliyunSnapshotRemote
func (cli *ZSClient) CreateAliyunSnapshotRemote(params param.CreateAliyunSnapshotRemoteParam) (*view.CreateAliyunSnapshotRemoteEventView, error) {
	resp := view.CreateAliyunSnapshotRemoteEventView{}
	if err := cli.Post("v1/hybrid/aliyun/snapshot", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
