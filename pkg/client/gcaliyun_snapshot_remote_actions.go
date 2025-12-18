// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GCAliyunSnapshotRemote 操作GCAliyunSnapshotRemote
func (cli *ZSClient) GCAliyunSnapshotRemote(params param.GCAliyunSnapshotRemoteParam) (*view.GCAliyunSnapshotRemoteEventView, error) {
	resp := view.GCAliyunSnapshotRemoteEventView{}
	if err := cli.Post("v1/hybrid/aliyun/snapshot/{dataCenterUuid}/gc", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

