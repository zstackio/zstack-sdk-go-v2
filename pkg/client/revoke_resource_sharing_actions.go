// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RevokeResourceSharing operates on RevokeResourceSharing
func (cli *ZSClient) RevokeResourceSharing(uuid string, params param.RevokeResourceSharingParam) (*view.RevokeResourceSharingEventView, error) {
	resp := view.RevokeResourceSharingEventView{}
	if err := cli.Put("v1/accounts/resources/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
