// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateAccount updates Account
func (cli *ZSClient) UpdateAccount(uuid string, params param.UpdateAccountParam) (*view.UpdateAccountEventView, error) {
	resp := view.UpdateAccountEventView{}
	if err := cli.Put("v1/accounts/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
