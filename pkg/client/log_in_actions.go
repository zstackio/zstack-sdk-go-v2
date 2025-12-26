// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// LogIn operates on LogIn
func (cli *ZSClient) LogIn(uuid string, params param.LogInParam) (*view.LogInView, error) {
	resp := view.LogInView{}
	if err := cli.Put("v1/login", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
