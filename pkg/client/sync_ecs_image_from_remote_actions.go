// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SyncEcsImageFromRemote 操作SyncEcsImageFromRemote
func (cli *ZSClient) SyncEcsImageFromRemote(params param.SyncEcsImageFromRemoteParam) (*view.SyncEcsImageFromRemoteEventView, error) {
	resp := view.SyncEcsImageFromRemoteEventView{}
	if err := cli.Post("v1/hybrid/aliyun/image/{dataCenterUuid}/sync", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

