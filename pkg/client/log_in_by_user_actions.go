// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// LogInByUser operates on LogInByUser
func (cli *ZSClient) LogInByUser(uuid string, params param.LogInByUserParam) (*view.LogInView, error) {
	resp := view.LogInView{}
	if err := cli.Put("v1/accounts/users/login", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
