// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateUserProxyConfig updates UserProxyConfig
func (cli *ZSClient) UpdateUserProxyConfig(uuid string, params param.UpdateUserProxyConfigParam) (*view.UpdateUserProxyConfigEventView, error) {
	resp := view.UpdateUserProxyConfigEventView{}
	if err := cli.Put("v1/user-proxy-configs/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
