// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateRole creates Role
func (cli *ZSClient) CreateRole(params param.CreateRoleParam) (*view.CreateRoleEventView, error) {
	resp := view.CreateRoleEventView{}
	if err := cli.Post("v1/identities/roles", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
