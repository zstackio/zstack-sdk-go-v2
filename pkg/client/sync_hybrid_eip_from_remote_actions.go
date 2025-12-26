// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SyncHybridEipFromRemote operates on SyncHybridEipFromRemote
func (cli *ZSClient) SyncHybridEipFromRemote(uuid string, params param.SyncHybridEipFromRemoteParam) (*view.SyncHybridEipFromRemoteEventView, error) {
	resp := view.SyncHybridEipFromRemoteEventView{}
	if err := cli.Put("v1/hybrid/eip/{dataCenterUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
