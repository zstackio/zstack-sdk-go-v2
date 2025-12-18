// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SyncAINginxConfiguration 操作SyncAINginxConfiguration
func (cli *ZSClient) SyncAINginxConfiguration(params param.SyncAINginxConfigurationParam) (*view.SyncAINginxConfigurationView, error) {
	resp := view.SyncAINginxConfigurationView{}
	if err := cli.Post("v1/ai/nginx/sync", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

