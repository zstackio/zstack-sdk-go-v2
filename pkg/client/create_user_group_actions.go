// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateUserGroup creates UserGroup
func (cli *ZSClient) CreateUserGroup(params param.CreateUserGroupParam) (*view.CreateUserGroupEventView, error) {
	resp := view.CreateUserGroupEventView{}
	if err := cli.Post("v1/accounts/groups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
