// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateUserGroup updates UserGroup
func (cli *ZSClient) UpdateUserGroup(uuid string, params param.UpdateUserGroupParam) (*view.UpdateUserGroupEventView, error) {
	resp := view.UpdateUserGroupEventView{}
	if err := cli.Put("v1/accounts/groups/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
