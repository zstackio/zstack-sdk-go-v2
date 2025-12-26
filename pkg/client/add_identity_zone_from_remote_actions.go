// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddIdentityZoneFromRemote adds IdentityZoneFromRemote
func (cli *ZSClient) AddIdentityZoneFromRemote(params param.AddIdentityZoneFromRemoteParam) (*view.AddIdentityZoneFromRemoteEventView, error) {
	resp := view.AddIdentityZoneFromRemoteEventView{}
	if err := cli.Post("v1/hybrid/identity-zone", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
