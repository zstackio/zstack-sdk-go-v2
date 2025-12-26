// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateFlkSecSecretResourcePool creates FlkSecSecretResourcePool
func (cli *ZSClient) CreateFlkSecSecretResourcePool(params param.CreateFlkSecSecretResourcePoolParam) (*view.CreateSecretResourcePoolEventView, error) {
	resp := view.CreateSecretResourcePoolEventView{}
	if err := cli.Post("v1/secret-resource-pool/flkSec", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
