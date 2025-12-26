// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateCSPSecretResourcePool creates CSPSecretResourcePool
func (cli *ZSClient) CreateCSPSecretResourcePool(params param.CreateCSPSecretResourcePoolParam) (*view.CreateSecretResourcePoolEventView, error) {
	resp := view.CreateSecretResourcePoolEventView{}
	if err := cli.Post("v1/secret-resource-pool/csp", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
