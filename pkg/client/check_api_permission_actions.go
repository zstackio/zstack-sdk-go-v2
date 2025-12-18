// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CheckApiPermission 操作CheckApiPermission
func (cli *ZSClient) CheckApiPermission(uuid string, params param.CheckApiPermissionParam) (*view.CheckApiPermissionView, error) {
	resp := view.CheckApiPermissionView{}
	if err := cli.Put("v1/accounts/permissions/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

