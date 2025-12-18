// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateCSPSecretResourcePool creates CSPSecretResourcePool
func (cli *ZSClient) CreateCSPSecretResourcePool(params param.CreateCSPSecretResourcePoolParam) (*view.CreateSecretResourcePoolEventView, error) {
	resp := view.CreateSecretResourcePoolEventView{}
	if err := cli.Post("v1/secret-resource-pool/csp", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
