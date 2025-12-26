// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SyncDiskFromAliyunFromRemote operates on SyncDiskFromAliyunFromRemote
func (cli *ZSClient) SyncDiskFromAliyunFromRemote(params param.SyncDiskFromAliyunFromRemoteParam) (*view.SyncDiskFromAliyunFromRemoteEventView, error) {
	resp := view.SyncDiskFromAliyunFromRemoteEventView{}
	if err := cli.Post("v1/hybrid/aliyun/disk/{identityUuid}/sync", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
