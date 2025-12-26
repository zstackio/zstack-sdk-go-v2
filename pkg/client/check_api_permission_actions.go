// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CheckApiPermission operates on CheckApiPermission
func (cli *ZSClient) CheckApiPermission(uuid string, params param.CheckApiPermissionParam) (*view.CheckApiPermissionView, error) {
	resp := view.CheckApiPermissionView{}
	if err := cli.Put("v1/accounts/permissions/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
