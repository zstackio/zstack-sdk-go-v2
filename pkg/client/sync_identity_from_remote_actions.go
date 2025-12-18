// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SyncIdentityFromRemote 操作SyncIdentityFromRemote
func (cli *ZSClient) SyncIdentityFromRemote(params param.SyncIdentityFromRemoteParam) (*view.SyncIdentityFromRemoteEventView, error) {
	var resp view.SyncIdentityFromRemoteEventView
	if err := cli.Get("v1/hybrid/identity-zone/{uuid}/sync", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

