// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateRole updates Role
func (cli *ZSClient) UpdateRole(uuid string, params param.UpdateRoleParam) (*view.UpdateRoleEventView, error) {
	resp := view.UpdateRoleEventView{}
	if err := cli.Put("v1/identities/roles/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
