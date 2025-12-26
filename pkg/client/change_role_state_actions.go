// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeRoleState changes RoleState
func (cli *ZSClient) ChangeRoleState(uuid string, params param.ChangeRoleStateParam) (*view.ChangeRoleStateEventView, error) {
	resp := view.ChangeRoleStateEventView{}
	if err := cli.Put("v1/identities/roles/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
