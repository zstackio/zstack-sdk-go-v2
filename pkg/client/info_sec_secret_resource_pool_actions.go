// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateInfoSecSecretResourcePool 创建InfoSecSecretResourcePool
func (cli *ZSClient) CreateInfoSecSecretResourcePool(params param.CreateInfoSecSecretResourcePoolParam) (*view.CreateSecretResourcePoolEventView, error) {
	resp := view.CreateSecretResourcePoolEventView{}
	if err := cli.Post("v1/secret-resource-pool/infoSec", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

