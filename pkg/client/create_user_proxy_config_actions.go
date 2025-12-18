// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateUserProxyConfig creates UserProxyConfig
func (cli *ZSClient) CreateUserProxyConfig(params param.CreateUserProxyConfigParam) (*view.CreateUserProxyConfigEventView, error) {
	resp := view.CreateUserProxyConfigEventView{}
	if err := cli.Post("v1/user-proxy-configs", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
