// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CheckResourcePermission operates on CheckResourcePermission
func (cli *ZSClient) CheckResourcePermission(params param.CheckResourcePermissionParam) (*view.CheckResourcePermissionView, error) {
	var resp view.CheckResourcePermissionView
	if err := cli.Get("v1/accounts/resource/api-permissions", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
