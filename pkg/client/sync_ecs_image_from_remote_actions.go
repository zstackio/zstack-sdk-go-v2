// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SyncEcsImageFromRemote operates on SyncEcsImageFromRemote
func (cli *ZSClient) SyncEcsImageFromRemote(params param.SyncEcsImageFromRemoteParam) (*view.SyncEcsImageFromRemoteEventView, error) {
	resp := view.SyncEcsImageFromRemoteEventView{}
	if err := cli.Post("v1/hybrid/aliyun/image/{dataCenterUuid}/sync", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
