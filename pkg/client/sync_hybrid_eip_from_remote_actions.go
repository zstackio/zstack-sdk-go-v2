// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SyncHybridEipFromRemote operates on SyncHybridEipFromRemote
func (cli *ZSClient) SyncHybridEipFromRemote(uuid string, params param.SyncHybridEipFromRemoteParam) (*view.SyncHybridEipFromRemoteEventView, error) {
	resp := view.SyncHybridEipFromRemoteEventView{}
	if err := cli.Put("v1/hybrid/eip/{dataCenterUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
