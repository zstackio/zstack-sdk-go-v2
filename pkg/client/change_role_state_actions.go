// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeRoleState changes RoleState
func (cli *ZSClient) ChangeRoleState(uuid string, params param.ChangeRoleStateParam) (*view.ChangeRoleStateEventView, error) {
	resp := view.ChangeRoleStateEventView{}
	if err := cli.Put("v1/identities/roles/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
