// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddIdentityZoneFromRemote adds IdentityZoneFromRemote
func (cli *ZSClient) AddIdentityZoneFromRemote(params param.AddIdentityZoneFromRemoteParam) (*view.AddIdentityZoneFromRemoteEventView, error) {
	resp := view.AddIdentityZoneFromRemoteEventView{}
	if err := cli.Post("v1/hybrid/identity-zone", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
