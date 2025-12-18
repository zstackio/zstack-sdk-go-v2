// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeResourceOwner changes ResourceOwner
func (cli *ZSClient) ChangeResourceOwner(uuid string, params param.ChangeResourceOwnerParam) (*view.ChangeResourceOwnerEventView, error) {
	resp := view.ChangeResourceOwnerEventView{}
	if err := cli.Put("v1/account/{accountUuid}/resources", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
