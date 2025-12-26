// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateUser creates User
func (cli *ZSClient) CreateUser(params param.CreateUserParam) (*view.CreateUserEventView, error) {
	resp := view.CreateUserEventView{}
	if err := cli.Post("v1/accounts/users", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
