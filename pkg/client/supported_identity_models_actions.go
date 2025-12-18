// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetSupportedIdentityModels 获取SupportedIdentityModels详情
func (cli *ZSClient) GetSupportedIdentityModels(uuid string) (*view.GetSupportedIdentityModelsView, error) {
	var resp view.GetSupportedIdentityModelsView
	if err := cli.Get("v1/identity-models", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

