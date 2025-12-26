// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateInfoSecSecretResourcePool creates InfoSecSecretResourcePool
func (cli *ZSClient) CreateInfoSecSecretResourcePool(params param.CreateInfoSecSecretResourcePoolParam) (*view.CreateSecretResourcePoolEventView, error) {
	resp := view.CreateSecretResourcePoolEventView{}
	if err := cli.Post("v1/secret-resource-pool/infoSec", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
