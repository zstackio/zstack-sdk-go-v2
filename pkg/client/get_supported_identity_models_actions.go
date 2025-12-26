// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetSupportedIdentityModels gets SupportedIdentityModels by uuid
func (cli *ZSClient) GetSupportedIdentityModels(uuid string) (*view.GetSupportedIdentityModelsView, error) {
	var resp view.GetSupportedIdentityModelsView
	if err := cli.Get("v1/identity-models", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
