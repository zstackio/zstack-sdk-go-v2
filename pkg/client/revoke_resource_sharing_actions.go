// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RevokeResourceSharing 操作RevokeResourceSharing
func (cli *ZSClient) RevokeResourceSharing(uuid string, params param.RevokeResourceSharingParam) (*view.RevokeResourceSharingEventView, error) {
	resp := view.RevokeResourceSharingEventView{}
	if err := cli.Put("v1/accounts/resources/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

