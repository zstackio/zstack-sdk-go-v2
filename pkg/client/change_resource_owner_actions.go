// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeResourceOwner changes ResourceOwner
func (cli *ZSClient) ChangeResourceOwner(uuid string, params param.ChangeResourceOwnerParam) (*view.ChangeResourceOwnerEventView, error) {
	resp := view.ChangeResourceOwnerEventView{}
	if err := cli.Put("v1/account/{accountUuid}/resources", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
