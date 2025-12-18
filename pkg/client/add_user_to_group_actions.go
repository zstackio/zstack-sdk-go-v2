// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddUserToGroup 操作AddUserToGroup
func (cli *ZSClient) AddUserToGroup(params param.AddUserToGroupParam) (*view.AddUserToGroupEventView, error) {
	resp := view.AddUserToGroupEventView{}
	if err := cli.Post("v1/accounts/groups/{groupUuid}/users", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

