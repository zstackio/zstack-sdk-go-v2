// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetIdentityZoneFromRemote gets IdentityZoneFromRemote by uuid
func (cli *ZSClient) GetIdentityZoneFromRemote(uuid string) (*view.GetIdentityZoneFromRemoteView, error) {
	var resp view.GetIdentityZoneFromRemoteView
	if err := cli.Get("v1/hybrid/identity-zone/remote", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
