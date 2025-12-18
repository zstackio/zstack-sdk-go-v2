// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateSecretResourcePool 更新SecretResourcePool
func (cli *ZSClient) UpdateSecretResourcePool(uuid string, params param.UpdateSecretResourcePoolParam) (*view.UpdateSecretResourcePoolEventView, error) {
	resp := view.UpdateSecretResourcePoolEventView{}
	if err := cli.Put("v1/secret-resource-pool/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

