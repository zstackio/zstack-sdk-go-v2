// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateHaiTaiSecretResourcePool 创建HaiTaiSecretResourcePool
func (cli *ZSClient) CreateHaiTaiSecretResourcePool(params param.CreateHaiTaiSecretResourcePoolParam) (*view.CreateSecretResourcePoolEventView, error) {
	resp := view.CreateSecretResourcePoolEventView{}
	if err := cli.Post("v1/secret-resource-pool/haitai", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

