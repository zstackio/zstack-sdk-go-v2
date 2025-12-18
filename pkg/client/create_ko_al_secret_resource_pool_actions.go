// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateKoAlSecretResourcePool creates KoAlSecretResourcePool
func (cli *ZSClient) CreateKoAlSecretResourcePool(params param.CreateKoAlSecretResourcePoolParam) (*view.CreateSecretResourcePoolEventView, error) {
	resp := view.CreateSecretResourcePoolEventView{}
	if err := cli.Post("v1/secret-resource-pool/koal", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
