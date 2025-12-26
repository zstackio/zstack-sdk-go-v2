// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// LogInByAccount operates on LogInByAccount
func (cli *ZSClient) LogInByAccount(uuid string, params param.LogInByAccountParam) (*view.LogInView, error) {
	resp := view.LogInView{}
	if err := cli.Put("v1/accounts/login", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
