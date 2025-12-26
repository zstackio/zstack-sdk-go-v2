// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SyncAliyunSnapshotRemote operates on SyncAliyunSnapshotRemote
func (cli *ZSClient) SyncAliyunSnapshotRemote(params param.SyncAliyunSnapshotRemoteParam) (*view.SyncAliyunSnapshotRemoteEventView, error) {
	resp := view.SyncAliyunSnapshotRemoteEventView{}
	if err := cli.Post("v1/hybrid/aliyun/snapshot/{dataCenterUuid}/sync", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
