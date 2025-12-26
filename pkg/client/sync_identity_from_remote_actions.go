// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SyncIdentityFromRemote operates on SyncIdentityFromRemote
func (cli *ZSClient) SyncIdentityFromRemote(params param.SyncIdentityFromRemoteParam) (*view.SyncIdentityFromRemoteEventView, error) {
	var resp view.SyncIdentityFromRemoteEventView
	if err := cli.Get("v1/hybrid/identity-zone/{uuid}/sync", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
