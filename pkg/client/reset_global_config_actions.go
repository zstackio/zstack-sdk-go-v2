// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ResetGlobalConfig 操作ResetGlobalConfig
func (cli *ZSClient) ResetGlobalConfig(uuid string, params param.ResetGlobalConfigParam) (*view.ResetGlobalConfigEventView, error) {
	resp := view.ResetGlobalConfigEventView{}
	if err := cli.Put("v1/global-configurations/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

