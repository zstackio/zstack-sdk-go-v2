// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateSanSecSecretResourcePool creates SanSecSecretResourcePool
func (cli *ZSClient) CreateSanSecSecretResourcePool(params param.CreateSanSecSecretResourcePoolParam) (*view.CreateSecretResourcePoolEventView, error) {
	resp := view.CreateSecretResourcePoolEventView{}
	if err := cli.Post("v1/secret-resource-pool/sanSec", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
