// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GCAliyunSnapshotRemote operates on GCAliyunSnapshotRemote
func (cli *ZSClient) GCAliyunSnapshotRemote(params param.GCAliyunSnapshotRemoteParam) (*view.GCAliyunSnapshotRemoteEventView, error) {
	resp := view.GCAliyunSnapshotRemoteEventView{}
	if err := cli.Post("v1/hybrid/aliyun/snapshot/{dataCenterUuid}/gc", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
