// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateUserProxyConfig creates UserProxyConfig
func (cli *ZSClient) CreateUserProxyConfig(params param.CreateUserProxyConfigParam) (*view.CreateUserProxyConfigEventView, error) {
	resp := view.CreateUserProxyConfigEventView{}
	if err := cli.Post("v1/user-proxy-configs", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
