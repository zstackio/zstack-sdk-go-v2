// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateAccount updates Account
func (cli *ZSClient) UpdateAccount(uuid string, params param.UpdateAccountParam) (*view.UpdateAccountEventView, error) {
	resp := view.UpdateAccountEventView{}
	if err := cli.Put("v1/accounts/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
