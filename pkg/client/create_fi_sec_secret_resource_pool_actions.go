// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateFiSecSecretResourcePool creates FiSecSecretResourcePool
func (cli *ZSClient) CreateFiSecSecretResourcePool(params param.CreateFiSecSecretResourcePoolParam) (*view.CreateSecretResourcePoolEventView, error) {
	resp := view.CreateSecretResourcePoolEventView{}
	if err := cli.Post("v1/secret-resource-pool/fiSec", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
