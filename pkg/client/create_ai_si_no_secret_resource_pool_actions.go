// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateAiSiNoSecretResourcePool creates AiSiNoSecretResourcePool
func (cli *ZSClient) CreateAiSiNoSecretResourcePool(params param.CreateAiSiNoSecretResourcePoolParam) (*view.CreateSecretResourcePoolEventView, error) {
	resp := view.CreateSecretResourcePoolEventView{}
	if err := cli.Post("v1/secret-resource-pool/aisino", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
