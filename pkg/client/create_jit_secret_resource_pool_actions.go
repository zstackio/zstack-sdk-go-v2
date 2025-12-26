// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateJitSecretResourcePool creates JitSecretResourcePool
func (cli *ZSClient) CreateJitSecretResourcePool(params param.CreateJitSecretResourcePoolParam) (*view.CreateSecretResourcePoolEventView, error) {
	resp := view.CreateSecretResourcePoolEventView{}
	if err := cli.Post("v1/secret-resource-pool/jit", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
