// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateAccount creates Account
func (cli *ZSClient) CreateAccount(params param.CreateAccountParam) (*view.CreateAccountEventView, error) {
	resp := view.CreateAccountEventView{}
	if err := cli.Post("v1/accounts", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
