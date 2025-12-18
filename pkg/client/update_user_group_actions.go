// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateUserGroup updates UserGroup
func (cli *ZSClient) UpdateUserGroup(uuid string, params param.UpdateUserGroupParam) (*view.UpdateUserGroupEventView, error) {
	resp := view.UpdateUserGroupEventView{}
	if err := cli.Put("v1/accounts/groups/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
