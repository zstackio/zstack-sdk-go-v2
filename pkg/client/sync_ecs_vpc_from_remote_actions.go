// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SyncEcsVpcFromRemote operates on SyncEcsVpcFromRemote
func (cli *ZSClient) SyncEcsVpcFromRemote(params param.SyncEcsVpcFromRemoteParam) (*view.SyncEcsVpcFromRemoteEventView, error) {
	resp := view.SyncEcsVpcFromRemoteEventView{}
	if err := cli.Post("v1/hybrid/aliyun/vpc/{dataCenterUuid}/sync", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
