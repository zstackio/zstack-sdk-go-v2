// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CheckResourcePermission 操作CheckResourcePermission
func (cli *ZSClient) CheckResourcePermission(params param.CheckResourcePermissionParam) (*view.CheckResourcePermissionView, error) {
	var resp view.CheckResourcePermissionView
	if err := cli.Get("v1/accounts/resource/api-permissions", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

