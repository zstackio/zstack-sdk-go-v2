// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateAiSiNoSecretResourcePool 创建AiSiNoSecretResourcePool
func (cli *ZSClient) CreateAiSiNoSecretResourcePool(params param.CreateAiSiNoSecretResourcePoolParam) (*view.CreateSecretResourcePoolEventView, error) {
	resp := view.CreateSecretResourcePoolEventView{}
	if err := cli.Post("v1/secret-resource-pool/aisino", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

