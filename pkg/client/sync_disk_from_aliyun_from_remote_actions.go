// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SyncDiskFromAliyunFromRemote 操作SyncDiskFromAliyunFromRemote
func (cli *ZSClient) SyncDiskFromAliyunFromRemote(params param.SyncDiskFromAliyunFromRemoteParam) (*view.SyncDiskFromAliyunFromRemoteEventView, error) {
	resp := view.SyncDiskFromAliyunFromRemoteEventView{}
	if err := cli.Post("v1/hybrid/aliyun/disk/{identityUuid}/sync", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

