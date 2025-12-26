// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddUserToGroup adds UserToGroup
func (cli *ZSClient) AddUserToGroup(params param.AddUserToGroupParam) (*view.AddUserToGroupEventView, error) {
	resp := view.AddUserToGroupEventView{}
	if err := cli.Post("v1/accounts/groups/{groupUuid}/users", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
