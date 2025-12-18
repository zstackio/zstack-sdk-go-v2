// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateUser creates User
func (cli *ZSClient) CreateUser(params param.CreateUserParam) (*view.CreateUserEventView, error) {
	resp := view.CreateUserEventView{}
	if err := cli.Post("v1/accounts/users", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
