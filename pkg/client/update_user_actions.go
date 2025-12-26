// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateUser updates User
func (cli *ZSClient) UpdateUser(uuid string, params param.UpdateUserParam) (*view.UpdateUserEventView, error) {
	resp := view.UpdateUserEventView{}
	if err := cli.Put("v1/accounts/users/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
