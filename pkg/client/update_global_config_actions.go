// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateGlobalConfig updates GlobalConfig
func (cli *ZSClient) UpdateGlobalConfig(uuid string, params param.UpdateGlobalConfigParam) (*view.UpdateGlobalConfigEventView, error) {
	resp := view.UpdateGlobalConfigEventView{}
	if err := cli.Put("v1/global-configurations/{category}/{name}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
