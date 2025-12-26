// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SyncEcsInstanceFromRemote operates on SyncEcsInstanceFromRemote
func (cli *ZSClient) SyncEcsInstanceFromRemote(params param.SyncEcsInstanceFromRemoteParam) (*view.SyncEcsInstanceFromRemoteEventView, error) {
	resp := view.SyncEcsInstanceFromRemoteEventView{}
	if err := cli.Post("v1/hybrid/aliyun/ecs/{dataCenterUuid}/sync", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
