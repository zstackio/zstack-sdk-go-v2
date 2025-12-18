// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SyncAliyunSnapshotRemote 操作SyncAliyunSnapshotRemote
func (cli *ZSClient) SyncAliyunSnapshotRemote(params param.SyncAliyunSnapshotRemoteParam) (*view.SyncAliyunSnapshotRemoteEventView, error) {
	resp := view.SyncAliyunSnapshotRemoteEventView{}
	if err := cli.Post("v1/hybrid/aliyun/snapshot/{dataCenterUuid}/sync", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

